<template>
  <div>
    <div class="detail-operation">
      <el-button icon="view" size="small" type="primary" plain @click="viewNode">查看</el-button>
      <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="nodeDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="nodeDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="nodeDetail.spec && nodeDetail.status" title="资源信息" name="spec">
          <div class="info-box">
            <div v-if="nodeDetail.spec.podCIDR" class="row">
              <div class="item">
                <p>Pod CIDR</p>
                <div class="content">
                  <span class="shadow">{{ nodeDetail.spec.podCIDR }}</span>
                </div>
              </div>
            </div>
            <div v-if="nodeDetail.status.addresses" class="row">
              <div class="item">
                <p>地址</p>
                <div class="content">
                  <span v-for="item in nodeDetail.status.addresses" :key="item.type" class="shadow">{{ item.type
                  }}: {{
                    item.address
                  }}</span>
                </div>
              </div>
            </div>
            <div v-if="nodeDetail.spec.taints" class="row">
              <div class="row">
                <div class="item">
                  <p>污点</p>
                  <div class="content">
                    <span v-for="item in nodeDetail.spec.taints" :key="item.type" class="shadow">{{ item.key }}={{
                      item.effect
                    }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="nodeDetail.status?.nodeInfo" title="系统信息" name="nodeInfo">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>操作系统</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.operatingSystem }}</span>
              </div>
              <div class="item">
                <p>操作系统镜像</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.osImage }}</span>
              </div>
              <div class="item">
                <p>架构</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.architecture }}</span>
              </div>
              <div class="item">
                <p>内核版本</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.kernelVersion }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>kubelet运行端口</p>
                <span class="content">{{ nodeDetail.status.Port.kubeletEndpoint.Port }}</span>
              </div>
              <div class="item">
                <p>kubelet版本</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.kubeletVersion }}</span>
              </div>
              <div class="item">
                <p>kube-proxy版本</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.kubeProxyVersion }}</span>
              </div>
              <div class="item">
                <p>容器runtime版本</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.containerRuntimeVersion }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>机器ID</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.machineID }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>系统UUID</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.systemUUID }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>启动ID</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.bootID }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="nodeDetail.allocatedResources" title="分配" name="allocated">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>CPU预留</p>
                <span class="content">{{ nodeDetail.allocatedResources.cpuRequests / 1000 }} cores ({{
                  nodeDetail.allocatedResources.cpuRequestsFraction.toFixed(1)
                }}%)</span>
              </div>
              <div class="item">
                <p>CPU限制</p>
                <span class="content">{{ nodeDetail.allocatedResources.cpuLimits / 1000 }} cores ({{
                  nodeDetail.allocatedResources.cpuLimitsFraction.toFixed(1)
                }}%) | Capacity {{
                  nodeDetail.allocatedResources.cpuCapacity / 1000
                }} cores</span>
              </div>
              <div class="item">
                <p>内存预留</p>
                <span class="content">{{ Math.round(nodeDetail.allocatedResources.memoryRequests / 1024 / 1024) }} MB
                  ({{ nodeDetail.allocatedResources.memoryRequestsFraction.toFixed(1) }}%)</span>
              </div>
              <div class="item">
                <p>内存限制</p>
                <span class="content">{{ Math.round(nodeDetail.allocatedResources.memoryLimits / 1024 / 1024) }} MB ({{
                  nodeDetail.allocatedResources.memoryLimitsFraction.toFixed(1)
                }}%) | Capacity {{
                  Math.round(nodeDetail.allocatedResources.memoryCapacity / 1024 / 1024)
                }} MB</span>
              </div>
              <div class="item">
                <p>Pods</p>
                <span class="content">{{ nodeDetail.allocatedResources.allocatedPods }} ({{
                  nodeDetail.allocatedResources.podFraction.toFixed(1)
                }}%) | Capacity {{
                  nodeDetail.allocatedResources.podCapacity
                }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="nodeDetail.status?.conditions" title="状态" name="status">
          <div class="info-table">
            <el-table :data="nodeDetail.status.conditions">
              <el-table-column prop="type" label="类别" min-width="120" />
              <el-table-column prop="status" label="状态" />
              <el-table-column label="最近检查时间" width="200">
                <template #default="scope">{{ formatDate(scope.row.lastHeartbeatTime) }}</template>
              </el-table-column>
              <el-table-column label="最近迁移时间" width="200">
                <template #default="scope">{{ formatDate(scope.row.lastTransitionTime) }}</template>
              </el-table-column>
              <el-table-column prop="reason" label="原因" min-width="180" />
              <el-table-column prop="message" label="信息" min-width="200" />
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item title="Pods" name="pods">
          <div class="info-table">
            <pod-brief :pods="nodePods" />
            <div class="gva-pagination">
              <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :total="total"
                layout="total, sizes, prev, pager, next, jumper"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
              />
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="nodeFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getNodeDetail, getNodeRaw, getNodePods } from '@/api/kubernetes/node'
import { statusPodFilter } from '@/mixin/filter.js'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import PodBrief from '@/components/kubernetes/pod/brief.vue'
import VueCodeMirror from '@/components/codeMirror/index.vue'
export default {
  name: 'NodeDetail',
  components: {
    MetaData,
    VueCodeMirror,
    PodBrief
  },
  setup() {
    const activeNames = ref(['metadata', 'spec', 'nodeInfo', 'allocated', 'status', 'pods'])
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const nodeDetail = ref({})
    const nodeFormat = ref({})
    const dialogFormVisible = ref(false)
    const nodePods = ref([])

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

    const getNodePodsData = async() => {
      const res = await getNodePods({ node_name: nodeName, page: page.value, pageSize: pageSize.value })
      if (res.code === 0) {
        nodePods.value = res.data.list
        total.value = res.data.total
        page.value = res.data.page
        pageSize.value = res.data.pageSize
      }
    }
    getNodePodsData()

    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getNodePodsData()
    }

    const handleCurrentChange = (val) => {
      page.value = val
      getNodePodsData()
    }

    // 操作
    const viewNode = async() => {
      const res = await getNodeRaw({ name: nodeName })
      if (res.code === 0) {
        nodeFormat.value = JSON.stringify(res.data)
      }
      dialogFormVisible.value = true
    }

    return {
      activeNames,
      nodeDetail,
      formatDate,
      nodeFormat,
      nodePods,
      dialogFormVisible,
      // filter
      statusPodFilter,
      // 函数
      viewNode,
      // 分页
      total,
      page,
      pageSize,
      handleSizeChange,
      handleCurrentChange
    }
  }
}
</script>
