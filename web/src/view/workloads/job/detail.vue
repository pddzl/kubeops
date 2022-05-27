<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewJob">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="jobDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="jobDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="jobDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>Parallelism</p>
                <span class="content">{{ jobDetail.spec.parallelism }}</span>
              </div>
              <div class="item">
                <p>Completions</p>
                <span class="content">{{ jobDetail.spec.completions }}</span>
              </div>
              <div class="item">
                <p>BackoffLimit</p>
                <span class="content">{{ jobDetail.spec.backoffLimit }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>Selector:</p>
                <div class="content">
                  <span v-for="(label, index) in jobDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="jobDetail.status" title="状态" name="status">
          <div v-if="jobDetail.status.conditions" class="info-table">
            <el-table :data="jobDetail.status.conditions">
              <el-table-column label="类别" prop="type" />
              <el-table-column label="状态" prop="status">
                <template #default="scope">
                  <el-tag :type="statusRsFilter(scope.row.status)" size="small">
                    {{ scope.row.status }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column label="上次探测时间">
                <template #default="scope">
                  {{ formatDate(scope.row.lastProbeTime) }}
                </template>
              </el-table-column>
              <el-table-column label="上次迁移时间">
                <template #default="scope">
                  {{ formatDate(scope.row.lastTransitionTime) }}
                </template>
              </el-table-column>
              <el-table-column label="原因">
                <template #default="scope">
                  <span v-if="scope.row.reason">{{ scope.row.reason }}</span>
                  <span v-else>-</span>
                </template>
              </el-table-column>
              <el-table-column label="信息" prop="message" :show-overflow-tooltip="true">
                <template #default="scope">
                  <span v-if="scope.row.message">{{ scope.row.message }}</span>
                  <span v-else>-</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item title="Pods" name="pods">
          <div class="info-table">
            <el-table :data="jobPods">
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
                  <el-tag :type="statusPodFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
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
      <vue-code-mirror v-model:modelValue="jobFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getJobDetail, getJobPods, getJobRaw, deleteJob } from '@/api/kubernetes/job'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { statusPodFilter, statusRsFilter } from '@/mixin/filter.js'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'JobDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'spec', 'status', 'pods'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const job = route.query.job

    // 加载job详情
    const jobDetail = ref({})

    const getData = async() => {
      await getJobDetail({ namespace: namespace, job: job }).then(response => {
        if (response.code === 0) {
          jobDetail.value = response.data
        }
      })
    }
    getData()

    // 加载关联pods
    const jobPods = ref([])
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)

    const getJobPodsData = async() => {
      const table = await getJobPods({ page: page.value, pageSize: pageSize.value, namespace: namespace, job: job })
      if (table.code === 0) {
        jobPods.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getJobPodsData()

    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getJobPodsData()
    }

    const handleCurrentChange = (val) => {
      page.value = val
      getJobPodsData()
    }

    // 编辑
    const jobFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewJob = async() => {
      const result = await getJobRaw({ job: job, namespace: namespace })
      if (result.code === 0) {
        jobFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该Job, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteJob({ namespace: namespace, job: job })
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '删除成功!'
            })
          }
        })
    }

    return {
      // 响应式数据
      activeNames,
      jobDetail,
      jobPods,
      formatDate,
      dialogFormVisible,
      namespace,
      // filter
      statusPodFilter,
      statusRsFilter,
      // 操作
      viewJob,
      jobFormat,
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
