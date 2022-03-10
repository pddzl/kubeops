<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="容器">
          <el-select v-model="searchInfo.container" clearable placeholder="请选择">
            <el-option v-for="item in containers" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="行数">
          <el-select v-model="searchInfo.lines" clearable placeholder="请选择">
            <el-option v-for="item in lines" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
          <el-button size="small" type="primary" plain icon="download" @click="donwloadLog">下载日志</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="log">
      <p v-for="(log, key) in podLogList" :key="key" style="color: rgba(0, 128, 0, 0.856)">
        <span style="color: rgba(255, 255, 0, 0.671);">{{ key + 1 }}</span>
        {{ log }}
      </p>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { getPodLog, getPodDetail } from '@/api/kubernetes/pod'
export default {
  name: 'PodLog',
  setup() {
    // 路由
    const route = useRoute()
    const pod = route.query.pod
    const namespace = route.query.namespace

    // 响应式数据
    const podLogList = ref([])
    const containers = ref([])
    const lines = ref([20, 50, 100, 200, 500])
    const searchInfo = reactive({ namespace: namespace, pod: pod, container: '', lines: 50 })

    // 加载pod数据
    const getPodLogData = async() => {
      const podData = await getPodLog({ ...searchInfo })
      if (podData.code === 0) {
        podLogList.value = podData.data.split('\n')
      }
    }
    getPodLogData()

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

    // 下载日志
    const donwloadLog = async() => {
      const podLog = await getPodLog({ ...searchInfo })
      if (podLog.code === 0) {
        const url = window.URL.createObjectURL(new Blob([podLog.data]))
        const a = document.createElement('a')
        a.style.display = 'none'
        a.href = url
        a.setAttribute('download', 'pod-log.txt')
        document.body.appendChild(a)
        a.click()
        window.URL.revokeObjectURL(a.href)
        document.body.removeChild(a)
      }
    }

    const onSubmit = () => {
      getPodLogData()
    }

    const onReset = () => {
      searchInfo.container = ''
      searchInfo.lines = 50
    }

    return {
      // 响应式数据
      containers,
      searchInfo,
      podLogList,
      lines,
      // 函数
      donwloadLog,
      onSubmit,
      onReset
    }
  }
}
</script>

<style lang="scss" scoped>
.log {
  padding: 24px;
  background-color: rgba(128, 128, 128, 0.253);
  border-radius: 2px;
}
</style>
