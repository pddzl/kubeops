<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="容器">
          <el-select v-model="searchInfo.container" clearable placeholder="请选择">
            <el-option v-for="item in containers" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="ArrowRight" @click="onSubmit">进 入</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- <div class="gva-table-box">
    </div> -->
    <div ref="terminalRef" />
  </div>
</template>

<script>
// 引入xterm css
import 'xterm/css/xterm.css'
// import 'xterm/dist/addons/fullscreen/fullscreen.css'
import { Terminal } from 'xterm'
// 自适应插件
import { FitAddon } from 'xterm-addon-fit'
// 全屏插件
// import * as fullscreen from 'xterm/lib/addons/fullscreen/fullscreen'
// web链接插件
// import * as webLinks from 'xterm/lib/addons/webLinks/webLinks'
// websocket插件
import { AttachAddon } from 'xterm-addon-attach'
import { ref, reactive, onBeforeUnmount } from 'vue'
import { useRoute } from 'vue-router'
import { getPodDetail } from '@/api/kubernetes/pod'
export default {
  name: 'PodTerminal',
  setup() {
    // 路由
    const route = useRoute()
    const pod = route.query.pod
    const namespace = route.query.namespace

    // 响应式数据
    const containers = ref([])
    const searchInfo = reactive({ namespace: namespace, pod: pod, container: '' })
    const terminalRef = ref(null)

    // 获取container
    const getPodContainer = async() => {
      const podDetail = await getPodDetail({ namespace: namespace, pod: pod })
      if (podDetail.code === 0) {
        podDetail.data.containers.forEach(element => {
          containers.value.push(element.name)
        })
      }
    }
    getPodContainer()

    // xterm
    let terminal = null
    let ws = null

    const wsOnOpen = () => {
      ws.onopen = () => {
        console.log('初始化terminal')
        initTerminal()
        // terminal.write('\r')
        // const msg = { op: 'stdin', data: 'export TERM=xterm && clear \r' }
        // ws.send(JSON.stringify(msg))
        // ws.send(msg)
      }
    }

    const wsOnMessage = () => {
      ws.onmessage = (event) => {
        // const msg = JSON.parse(event.data)
        // if (msg.op === 'stdout') {
        //   terminal.write(msg.data)
        // } else {
        //   console.log('invalid msg op: ' + msg)
        // }
      }
    }

    const wsOnError = () => {
      ws.onerror = (error) => {
        console.log('[error] Connection error')
        terminal.write('error: ' + error.message)
        terminal.destroy()
      }
    }

    const wsOnClose = () => {
      ws.onclose = (event) => {
        if (event.wasClean) {
          console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`)
        } else {
          console.log('[close] Connection died')
          terminal.writeln('')
        }
        terminal.write('Connection Reset By Peer! Try Refresh.')
      }
    }

    const initWs = () => {
      // 打开websocket
      console.log('初始化websocket')
      ws = new WebSocket(`ws://10.192.104.101:8888/pod/getPodTerminal?namespace=${namespace}&pod=${pod}&container=${searchInfo.container}`)
      wsOnClose()
      if (terminal === null) {
        wsOnOpen()
      }
      wsOnMessage()
      wsOnError()
    }
    initWs()

    // 初始化终端
    const initTerminal = () => {
      terminal = new Terminal({
        rendererType: 'canvas',
        rows: 32,
        // cols: 40,
        convertEol: true, // 启用时，光标将设置为下一行的开头
        fontSize: 14,
        cursorBlink: true,
        disableStdin: false
      })
      const attachAddon = new AttachAddon(ws)
      const fitAddon = new FitAddon()
      terminal.loadAddon(attachAddon)
      terminal.loadAddon(fitAddon)
      terminal.open(terminalRef.value)
      fitAddon.fit()
      // console.log(terminal.cols, terminal.rows)
      terminal.focus()
      // terminal.onData(function(data) {
      //   const msg = { op: 'stdin', data: data }
      //   ws.send(JSON.stringify(msg))
      // })
      // terminal.toggleFullScreen(true)
      // terminal.on('data', function(data) {
      //   const msg = { op: 'stdin', data: data }
      //   ws.send(JSON.stringify(msg))
      // })
      // terminal.on('resize', function(size) {
      //   console.log('resize: ' + size)
      //   const msg = { op: 'resize', cols: size.cols, rows: rows }
      //   ws.send(JSON.stringify(msg))
      // })
    }

    const onSubmit = () => {
      if (searchInfo.container !== '') {
        if (Object.prototype.hasOwnProperty.call(ws, 'close')) {
          ws.close()
        }
        initWs()
      }
    }

    onBeforeUnmount(() => {
      wsOnClose()
    })

    return {
      // 响应式数据
      containers,
      searchInfo,
      terminalRef,
      // 函数
      onSubmit,
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
