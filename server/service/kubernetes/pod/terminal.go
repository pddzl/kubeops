package pod

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"

	"github.com/pddzl/kubeops/server/global"
)

const END_OF_TRANSMISSION = "\u0004"

// PtyHandler is what remotecommand expects from a pty
type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

type TerminalSession struct {
	wsConn   *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

func (p *PodService) NewTerminalSession(conn *websocket.Conn) (*TerminalSession, error) {
	session := &TerminalSession{
		wsConn:   conn,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}
	return session, nil
}

// TerminalMessage is the messaging protocol between ShellController and TerminalSession.
//
// OP      DIRECTION  FIELD(S) USED  DESCRIPTION
// ---------------------------------------------------------------------
// bind    fe->be     SessionID      Id sent back from TerminalResponse
// stdin   fe->be     Data           Keystrokes/paste buffer
// resize  fe->be     Rows, Cols     New terminal size
// stdout  be->fe     Data           Output from the process
// toast   be->fe     Data           OOB message to be shown to the user
type TerminalMessage struct {
	Op, Data, SessionID string
	Rows, Cols          uint16
}

// Next TerminalSize handles pty->process resize events
// Called in a loop from remotecommand as long as the process is running
func (t *TerminalSession) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}

// Read handles pty->process messages (stdin, resize)
// Called in a loop from remotecommand as long as the process is running
func (t *TerminalSession) Read(p []byte) (int, error) {
	_, message, err := t.wsConn.ReadMessage()
	global.KOP_LOG.Info("message", zap.String("ms", string(message)))
	if err != nil {
		return copy(p, END_OF_TRANSMISSION), err
	}
	var msg TerminalMessage
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		global.KOP_LOG.Error("json", zap.Error(err))
		return copy(p, END_OF_TRANSMISSION), err
	}
	switch msg.Op {
	case "stdin":
		global.KOP_LOG.Info("stdin debug")
		return copy(p, msg.Data), nil
	case "resize":
		t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	default:
		global.KOP_LOG.Error(fmt.Sprintf("unknown message type '%s'", msg.Op))
		return copy(p, END_OF_TRANSMISSION), fmt.Errorf("unknown message type '%s'", msg.Op)
	}
}

// Write called from remotecommand whenever there is any output
func (t *TerminalSession) Write(p []byte) (int, error) {
	msg, err := json.Marshal(TerminalMessage{
		Op:   "stdout",
		Data: string(p),
	})
	if err != nil {
		return 0, err
	}
	global.KOP_LOG.Info(string(msg))
	if err := t.wsConn.WriteMessage(websocket.TextMessage, p); err != nil {
		return 0, err
	}
	return len(p), nil
}

// Close session
func (t *TerminalSession) Close() error {
	close(t.doneChan)
	return t.wsConn.Close()
}

func (p *PodService) GetPodTerminal(namespace string, podName string, containerName string, cfg *rest.Config, ptyHandler PtyHandler) error {
	req := global.KOP_KUBERNETES.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Container: containerName,
		Command:   []string{"sh", "-c", "test -f /bin/bash && bash || sh"},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(cfg, "POST", req.URL())
	if err != nil {
		return err
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:             ptyHandler,
		Stdout:            ptyHandler,
		Stderr:            ptyHandler,
		TerminalSizeQueue: ptyHandler,
		Tty:               true,
	})
	if err != nil {
		return err
	}

	return nil
}
