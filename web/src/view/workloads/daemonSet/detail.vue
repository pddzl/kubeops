<template>
  <div>
    <div class="detail-operation">
      <el-button icon="view" size="small" type="primary" plain @click="viewDaemonSet">查看</el-button>
      <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="daemonSetDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="daemonSetDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="daemonSetDetail.spec" title="资源信息" name="resource">
          <div class="info-box">
            <div class="row">
              <div v-if="daemonSetDetail.spec.selector" class="item">
                <p>Selector</p>
                <div class="content">
                  <span v-for="(label, index) in daemonSetDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
            <div v-if="daemonSetDetail.spec.updateStrategy" class="row">
              <div class="item">
                <p>Strategy</p>
                <span class="content">{{ daemonSetDetail.spec.updateStrategy.type }}</span>
              </div>
              <div class="item">
                <p>maxUnavailable</p>
                <span class="content">{{ daemonSetDetail.spec.updateStrategy.rollingUpdate.maxUnavailable }}</span>
              </div>
              <div v-if="daemonSetDetail.spec.updateStrategy.type === 'RollingUpdate'" class="item">
                <p>maxSurge</p>
                <span class="content">{{ daemonSetDetail.spec.updateStrategy.rollingUpdate.maxUnavailable }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="daemonSetDetail.status" title="状态" name="status">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>currentNumberScheduled</p>
                <span class="content">{{ daemonSetDetail.status.currentNumberScheduled }}</span>
              </div>
              <div class="item">
                <p>desiredNumberScheduled</p>
                <span class="content">{{ daemonSetDetail.status.desiredNumberScheduled }}</span>
              </div>
              <div class="item">
                <p>numberReady</p>
                <span class="content">{{ daemonSetDetail.status.numberReady }}</span>
              </div>
              <div class="item">
                <p>numberAvailable</p>
                <span class="content">{{ daemonSetDetail.status.numberAvailable }}</span>
              </div>
            </div>
          </div>
          <div v-if="daemonSetDetail.status.conditions" style="margin-top: 20px; margin-right: 20px;">
            <el-table :data="daemonSetDetail.status.conditions">
              <el-table-column label="类别" prop="type" />
              <el-table-column label="状态" prop="status">
                <template #default="scope">
                  <el-tag :type="statusRsFilter(scope.row.status)" size="small">
                    {{ scope.row.status }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column label="上次迁移时间">
                <template #default="scope">
                  {{ formatDate(scope.row.lastTransitionTime) }}
                </template>
              </el-table-column>
              <el-table-column label="原因" prop="reason" />
              <el-table-column label="信息" prop="message" :show-overflow-tooltip="true" />
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item title="Pods" name="pods">
          <div style="padding-right: 20px;">
            <el-table :data="daemonSetPods">
              <el-table-column label="名称" prop="metadata.name" min-width="120">
                <template #default="scope">
                  <router-link
                    :to="{ name: 'pod_detail', query: { pod: scope.row.metadata.name, namespace: scope.row.metadata.namespace } }"
                  >
                    <el-link type="primary" :underline="false">{{ scope.row.metadata.name }}</el-link>
                  </router-link>
                </template>
              </el-table-column>
              <el-table-column label="命名空间" prop="metadata.namespace" />
              <el-table-column label="节点" prop="nodeName" />
              <el-table-column label="状态">
                <template #default="scope">
                  <el-tag
                    :type="statusPodFilter(scope.row.status)"
                    size="small"
                  >{{ scope.row.status }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="创建时间">
                <template #default="scope">{{ formatDate(scope.row.metadata.creationTimestamp) }}</template>
              </el-table-column>
            </el-table>
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
      <vue-code-mirror v-model:modelValue="daemonSetFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getDaemonSetDetail, getDaemonSetRaw, getDaemonSetPods, deleteDaemonSet } from '@/api/kubernetes/daemonSet'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { formatDate } from '@/utils/format'
import { statusRsFilter } from '@/mixin/filter.js'
import { statusPodFilter } from '@/mixin/filter.js'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'DaemonSetDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'resource', 'status', 'pods'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const daemonSet = route.query.daemonSet

    // 获取daemonSet详情
    const daemonSetDetail = ref({})

    const getDaemonSetDetailData = async() => {
      await getDaemonSetDetail({ namespace: namespace, daemonSet: daemonSet }).then(response => {
        if (response.code === 0) {
          daemonSetDetail.value = response.data
        }
      })
    }
    getDaemonSetDetailData()

    // 加载关联pods
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const daemonSetPods = ref([])

    const getDaemonSetPodsData = async() => {
      const table = await getDaemonSetPods({ page: page.value, pageSize: pageSize.value, namespace: namespace, daemonSet: daemonSet })
      if (table.code === 0) {
        daemonSetPods.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getDaemonSetPodsData()

    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getDaemonSetPodsData()
    }

    const handleCurrentChange = (val) => {
      page.value = val
      getDaemonSetPodsData()
    }

    // 编辑
    const daemonSetFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewDaemonSet = async() => {
      const result = await getDaemonSetRaw({ daemonSet: daemonSet, namespace: namespace })
      if (result.code === 0) {
        daemonSetFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该DaemonSet, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteDaemonSet({ namespace: namespace, daemonSet: daemonSet })
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '删除成功!'
            })
          }
        })
    }

    return {
      // 折叠面板
      activeNames,
      // data、查询相关
      namespace,
      daemonSetDetail,
      daemonSetPods,
      // filter
      statusRsFilter,
      statusPodFilter,
      // 格式化日期
      formatDate,
      // 编辑
      viewDaemonSet,
      daemonSetFormat,
      dialogFormVisible,
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange,
      // 删除
      deleteFunc
    }
  }
}
</script>

