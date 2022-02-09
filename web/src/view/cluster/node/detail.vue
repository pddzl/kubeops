<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="nodeDetail.objectMeta" class="object_meta" title="元数据" name="1">
        <span style="font-size: 14px; margin-right: 150px;">名称: {{ nodeDetail.objectMeta.name }}</span>
        <span style="font-size: 14px; margin-right: 150px;">UID: {{ nodeDetail.objectMeta.uid }}</span>
        <span style="font-size: 14px">创建时间: {{ formatDate(nodeDetail.objectMeta.createTimestamp) }}</span>
        <p>标签:</p>
        <div v-for="(label, index) in nodeDetail.objectMeta.labels" :key="index">
          <span class="label">{{ index }}<span v-if="label">:</span> {{ label }}</span>
        </div>
        <p>注释:</p>
        <div v-for="(label, index) in nodeDetail.objectMeta.annotations" :key="index">
          <span class="label">{{ index }}<span v-if="label">:</span> {{ label }}</span>
        </div>
      </el-collapse-item>
      <el-collapse-item class="resource" title="资源信息" name="2">
        <div>
          <p>Pod CIDR:</p>
          <span>{{ nodeDetail.podCIDR }}</span>
        </div>
        <div>
          <p>地址:</p>
          <span v-for="item in nodeDetail.addresses" :key="item.type">
            {{ item.type }}: {{ item.address }}
          </span>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="nodeDetail.nodeInfo" title="系统信息" name="3">
        <el-descriptions direction="vertical">
          <el-descriptions-item label="机器ID">{{ nodeDetail.nodeInfo.machineID }}</el-descriptions-item>
          <el-descriptions-item label="系统UUID">{{ nodeDetail.nodeInfo.systemUUID }}</el-descriptions-item>
          <el-descriptions-item label="启动ID">{{ nodeDetail.nodeInfo.bootID }}</el-descriptions-item>
        </el-descriptions>
        <el-descriptions direction="vertical" :column="7">
          <el-descriptions-item label="内核版本">{{ nodeDetail.nodeInfo.kernelVersion }}</el-descriptions-item>
          <el-descriptions-item label="操作系统镜像">{{ nodeDetail.nodeInfo.osImage }}</el-descriptions-item>
          <el-descriptions-item
            label="容器runtime版本"
          >{{ nodeDetail.nodeInfo.containerRuntimeVersion }}</el-descriptions-item>
          <el-descriptions-item label="kubelet版本">{{ nodeDetail.nodeInfo.kubeletVersion }}</el-descriptions-item>
          <el-descriptions-item label="kube-proxy版本">{{ nodeDetail.nodeInfo.kubeProxyVersion }}</el-descriptions-item>
          <el-descriptions-item label="操作系统">{{ nodeDetail.nodeInfo.operatingSystem }}</el-descriptions-item>
          <el-descriptions-item label="架构">{{ nodeDetail.nodeInfo.architecture }}</el-descriptions-item>
        </el-descriptions>
      </el-collapse-item>
      <el-collapse-item
        v-if="nodeDetail.allocatedResources"
        title="分配"
        name="4"
      >
        <el-descriptions direction="vertical" :column="5">
          <el-descriptions-item
            label="CPU预留"
          >{{ nodeDetail.allocatedResources.cpuRequests / 1000 }} cores ({{ nodeDetail.allocatedResources.cpuRequestsFraction.toFixed(1) }}%)</el-descriptions-item>
          <el-descriptions-item
            label="CPU限制"
          >{{ nodeDetail.allocatedResources.cpuLimits / 1000 }} cores ({{ nodeDetail.allocatedResources.cpuLimitsFraction.toFixed(1) }}%) | Capacity {{ nodeDetail.allocatedResources.cpuCapacity / 1000 }} cores</el-descriptions-item>
          <el-descriptions-item
            label="内存预留"
          >{{ Math.round(nodeDetail.allocatedResources.memoryRequests / 1024 / 1024) }} MB ({{ nodeDetail.allocatedResources.memoryRequestsFraction.toFixed(1) }}%)</el-descriptions-item>
          <el-descriptions-item
            label="内存限制"
          >{{ Math.round(nodeDetail.allocatedResources.memoryLimits / 1024 / 1024) }} MB ({{ nodeDetail.allocatedResources.memoryLimitsFraction.toFixed(1) }}%) | Capacity {{ Math.round(nodeDetail.allocatedResources.memoryCapacity / 1024 / 1024) }} MB</el-descriptions-item>
          <el-descriptions-item
            label="Pods"
          >{{ nodeDetail.allocatedResources.allocatedPods }} ({{ nodeDetail.allocatedResources.podFraction.toFixed(1) }}%) | Capacity {{ nodeDetail.allocatedResources.podCapacity }}</el-descriptions-item>
        </el-descriptions>
      </el-collapse-item>
      <el-collapse-item v-if="nodeDetail.conditions" title="状态" name="5">
        <el-table :data="nodeDetail.conditions">
          <el-table-column prop="type" label="类别" min-width="120" />
          <el-table-column prop="status" label="状态" />
          <el-table-column prop="reason" label="原因" min-width="180" />
          <el-table-column prop="message" label="信息" min-width="200" />
          <el-table-column label="最后检查时间" width="200">
            <template #default="scope">{{ formatDate(scope.row.lastProbeTime) }}</template>
          </el-table-column>
        </el-table>
      </el-collapse-item>
      <el-collapse-item v-if="nodeDetail.podList" title="Pods" name="6">
        <el-table :data="nodeDetail.podList">
          <el-table-column prop="name" label="名称" min-width="240" />
          <el-table-column prop="status" label="状态" min-width="110" />
          <el-table-column prop="namespace" label="命名空间" min-width="120" />
          <el-table-column prop="resource.cpuLimit" label="CPU预留" width="100" />
          <el-table-column prop="resource.memoryLimit" label="CPU限制" width="100" />
          <el-table-column prop="resource.cpuRequests" label="内存预留" width="100" />
          <el-table-column prop="resource.memoryRequests" label="内存限制" width="100" />
          <!-- <el-table-column prop="image" label="镜像" /> -->
          <el-table-column label="创建时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.createTimestamp) }}</template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="200">
            <template #default>
              <el-button icon="tickets" type="text" size="small">日志</el-button>
              <el-button icon="ArrowRight" type="text" size="small">终端</el-button>
              <el-button icon="delete" type="text" size="small">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getNodeDetail } from '@/api/kubernetes/node'
import { formatDate } from '@/utils/format'
export default {
  name: 'Detail',
  setup() {
    const activeNames = ref(['1', '2', '3', '4', '5', '6'])
    const nodeDetail = ref({})

    const route = useRoute()
    const nodeName = route.query.name

    // 加载node详情
    const getData = async() => {
      await getNodeDetail({ name: nodeName }).then(response => {
        if (response.code === 0) {
          // console.log('data', response.data)
          nodeDetail.value = response.data
        }
      })
    }
    getData()

    return {
      activeNames,
      nodeDetail,
      formatDate
    }
  }
}
</script>

<style scoped lang="scss">
@mixin shadow {
  background-color: rgba(128, 128, 128, 0.096);
  font-size: 13px;
  padding: 4px;
  border-radius: 8px;
}
.object_meta {
  p {
    font-size: 14px;
    margin-top: 15px;
  }
  .label {
    @include shadow;
  }
}
.resource {
  p {
    font-size: 14px;
  }
  span {
    @include shadow;
  }
  span:last-child {
    margin-left: 5px;
  }
}
.el-collapse {
  --el-collapse-header-font-size: 15px;
  .el-collapse-item {
    background-color: #ffffff;
    padding-left: 20px;
  }
  .el-collapse-item:not(:last-child) {
    margin-bottom: 15px;
  }
}
</style>
