<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="nodeDetail.objectMeta" title="元数据" name="1">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>名称</p>
              <span class="content">{{ nodeDetail.objectMeta.name }}</span>
            </div>
            <div>
              <p>UID</p>
              <span class="content">{{ nodeDetail.objectMeta.uid }}</span>
            </div>
            <div>
              <p>创建时间</p>
              <span class="content">{{ formatDate(nodeDetail.objectMeta.creationTimestamp) }}</span>
            </div>
          </div>
        </div>
        <div class="common_show">
          <p>标签:</p>
          <div v-for="(label, index) in nodeDetail.objectMeta.labels" :key="index">
            <span class="span-shadow">
              {{ index }}
              <span v-if="label">:</span>
              {{ label }}
            </span>
          </div>
        </div>
        <div class="common_show">
          <p>注释:</p>
          <div v-for="(label, index) in nodeDetail.objectMeta.annotations" :key="index">
            <span class="span-shadow">
              {{ index }}
              <span v-if="label">:</span>
              {{ label }}
            </span>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item title="资源信息" name="2">
        <div class="common_show">
          <p>Pod CIDR</p>
          <span class="span-shadow">{{ nodeDetail.podCIDR }}</span>
        </div>
        <div v-if="nodeDetail.addresses" class="common_show interval">
          <p>地址</p>
          <span
            v-for="item in nodeDetail.addresses"
            :key="item.type"
            class="span-shadow"
          >{{ item.type }}: {{ item.address }}</span>
        </div>
        <div v-if="nodeDetail.taints" class="common_show interval">
          <p>污点</p>
          <span
            v-for="item in nodeDetail.taints"
            :key="item.type"
            class="span-shadow"
          >{{ item.key }}={{ item.effect }}</span>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="nodeDetail.nodeInfo" title="系统信息" name="4">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>机器ID</p>
              <span class="content">{{ nodeDetail.nodeInfo.machineID }}</span>
            </div>
            <div>
              <p>系统UUID</p>
              <span class="content">{{ nodeDetail.nodeInfo.systemUUID }}</span>
            </div>
            <div>
              <p>启动ID</p>
              <span class="content">{{ nodeDetail.nodeInfo.bootID }}</span>
            </div>
            <div>
              <p>操作系统镜像</p>
              <span class="content">{{ nodeDetail.nodeInfo.osImage }}</span>
            </div>
            <div>
              <p>内核版本</p>
              <span class="content">{{ nodeDetail.nodeInfo.kernelVersion }}</span>
            </div>
            <div>
              <p>容器runtime版本</p>
              <span class="content">{{ nodeDetail.nodeInfo.containerRuntimeVersion }}</span>
            </div>
            <div>
              <p>kubelet版本</p>
              <span class="content">{{ nodeDetail.nodeInfo.kubeletVersion }}</span>
            </div>
            <div>
              <p>kube-proxy版本</p>
              <span class="content">{{ nodeDetail.nodeInfo.kubeProxyVersion }}</span>
            </div>
            <div>
              <p>操作系统</p>
              <span class="content">{{ nodeDetail.nodeInfo.operatingSystem }}</span>
            </div>
            <div>
              <p>架构</p>
              <span class="content">{{ nodeDetail.nodeInfo.architecture }}</span>
            </div>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="nodeDetail.allocatedResources" title="分配" name="4">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>CPU预留</p>
              <span
                class="content"
              >{{ nodeDetail.allocatedResources.cpuRequests / 1000 }} cores ({{ nodeDetail.allocatedResources.cpuRequestsFraction.toFixed(1) }}%)</span>
            </div>
            <div>
              <p>CPU限制</p>
              <span
                class="content"
              >{{ nodeDetail.allocatedResources.cpuLimits / 1000 }} cores ({{ nodeDetail.allocatedResources.cpuLimitsFraction.toFixed(1) }}%) | Capacity {{ nodeDetail.allocatedResources.cpuCapacity / 1000 }} cores</span>
            </div>
            <div>
              <p>内存预留</p>
              <span
                class="content"
              >{{ Math.round(nodeDetail.allocatedResources.memoryRequests / 1024 / 1024) }} MB ({{ nodeDetail.allocatedResources.memoryRequestsFraction.toFixed(1) }}%)</span>
            </div>
            <div>
              <p>内存限制</p>
              <span
                class="content"
              >{{ Math.round(nodeDetail.allocatedResources.memoryLimits / 1024 / 1024) }} MB ({{ nodeDetail.allocatedResources.memoryLimitsFraction.toFixed(1) }}%) | Capacity {{ Math.round(nodeDetail.allocatedResources.memoryCapacity / 1024 / 1024) }} MB</span>
            </div>
            <div>
              <p>Pods</p>
              <span
                class="content"
              >{{ nodeDetail.allocatedResources.allocatedPods }} ({{ nodeDetail.allocatedResources.podFraction.toFixed(1) }}%) | Capacity {{ nodeDetail.allocatedResources.podCapacity }}</span>
            </div>
          </div>
        </div>
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
          <el-table-column label="名称" min-width="200">
            <template #default="scope">
              <router-link
                :to="{ name: 'pod_detail', query: { pod: scope.row.name, namespace: scope.row.namespace } }"
              >
                <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
              </router-link>
            </template>
          </el-table-column>
          <el-table-column prop="status" label="状态" min-width="90">
            <template #default="scope">
              <el-tag :type="statusPodFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="namespace" label="命名空间" min-width="100" />
          <el-table-column prop="image" label="镜像" min-width="340" />
          <el-table-column label="创建时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
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
import { statusPodFilter } from '@/mixin/filter.js'
import { formatDate } from '@/utils/format'
export default {
  name: 'NodeDetail',
  setup() {
    const activeNames = ref(['1', '2', '3', '4', '5', '6'])
    const nodeDetail = ref({})

    const route = useRoute()
    const nodeName = route.query.name

    // 加载node详情
    const getData = async() => {
      await getNodeDetail({ name: nodeName }).then(response => {
        if (response.code === 0) {
          nodeDetail.value = response.data
        }
      })
    }
    getData()

    return {
      activeNames,
      nodeDetail,
      formatDate,
      // filter
      statusPodFilter
    }
  }
}
</script>

<style scoped lang="scss">
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
.interval {
  span:not(:last-child) {
    margin-right: 5px;
  }
}
</style>
