<template>
  <div>
    <div class="detail-operation">
      <el-button icon="view" size="small" type="primary" plain @click="editServicesSet">查看</el-button>
      <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="servicesDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="servicesDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="servicesDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div v-if="servicesDetail.spec.type" class="item">
                <p>类别</p>
                <span class="content">{{ servicesDetail.spec.type }}</span>
              </div>
              <div v-if="servicesDetail.spec.clusterIP" class="item">
                <p>集群IP</p>
                <span class="content">{{ servicesDetail.spec.clusterIP }}</span>
              </div>
              <div v-if="servicesDetail.spec.sessionAffinity" class="item">
                <p>sessionAffinity</p>
                <span class="content">{{ servicesDetail.spec.sessionAffinity }}</span>
              </div>
            </div>
            <div class="row">
              <div v-if="servicesDetail.spec.selector" class="item">
                <p>Selector</p>
                <div class="content">
                  <span v-for="(label, index) in servicesDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="servicesDetail.spec" title="端口映射" name="ports">
          <div class="info-table">
            <el-table :data="servicesDetail.spec.ports">
              <el-table-column label="名称" prop="name" />
              <el-table-column label="协议" prop="protocol" />
              <el-table-column label="port" prop="port" />
              <el-table-column label="targetPort" prop="targetPort" />
              <el-table-column v-if="servicesDetail.spec.type === 'NodePort'" label="nodePort" prop="nodePort" />
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item title="Pods" name="pods">
          <div class="info-table">
            <el-table :data="servicesPods">
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
      <vue-code-mirror v-model:modelValue="servicesFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getServicesDetail, getServicesRaw, getServicesPods, deleteServices } from '@/api/kubernetes/services'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import { statusRsFilter, statusPodFilter } from '@/mixin/filter.js'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'ServicesDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'spec', 'ports', 'pods'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const service = route.query.service

    // 获取daemonSet详情
    const servicesDetail = ref({})

    const getServicesDetailData = async() => {
      await getServicesDetail({ namespace: namespace, service: service }).then(response => {
        if (response.code === 0) {
          servicesDetail.value = response.data
        }
      })
    }
    getServicesDetailData()

    // 加载关联pods
    const servicesPods = ref([])
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)

    const getServicesPodsData = async() => {
      const table = await getServicesPods({ page: page.value, pageSize: pageSize.value, namespace: namespace, service: service })
      if (table.code === 0) {
        servicesPods.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getServicesPodsData()

    // 操作
    const servicesFormat = ref({})
    const dialogFormVisible = ref(false)

    const editServicesSet = async() => {
      const result = await getServicesRaw({ service: service, namespace: namespace })
      if (result.code === 0) {
        servicesFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getServicesPodsData()
    }

    const handleCurrentChange = (val) => {
      page.value = val
      getServicesPodsData()
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该Service, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteServices({ namespace: namespace, service: service })
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
      // detail
      servicesDetail,
      servicesPods,
      // filter
      statusRsFilter,
      statusPodFilter,
      // 格式化日期
      formatDate,
      // 操作
      servicesFormat,
      dialogFormVisible,
      editServicesSet,
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
