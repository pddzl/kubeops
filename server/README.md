## 编译

sudo env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o kop main.go