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
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div ref="terminalRef" />
    </div>
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
      }
    }

    const wsOnError = () => {
      ws.onerror = () => {
        console.log('websocket连接失败')
      }
    }

    const wsOnClose = () => {
      ws.onclose = (e) => {
        console.log('关闭websocket')
        console.log('websocket 断开: ' + e.code + ' ' + e.reason + ' ' + e.wasClean)
        console.log(e)
      }
    }

    const initWs = () => {
      // 打开websocket
      ws = new WebSocket(`ws://10.192.104.101:8888/pod/getPodTerminal?namespace=${namespace}&pod=${pod}`)
      wsOnClose()
      wsOnOpen()
      wsOnError()
    }
    initWs()

    // 初始化终端
    const initTerminal = () => {
      terminal = new Terminal({
        rendererType: 'canvas',
        rows: 40,
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
      terminal.focus()
    }

    const onSubmit = () => {
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
