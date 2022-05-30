<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true">
        <el-form-item label="命名空间">
          <el-select v-model="searchInfo.namespace" clearable placeholder="请选择">
            <el-option v-for="item in namespace" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column label="名称" min-width="220">
          <template #default="scope">
            <router-link
              :to="{ name: 'services_detail', query: { service: scope.row.name, namespace: scope.row.namespace } }"
            >
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="命名空间" prop="namespace" min-width="120" />
        <el-table-column label="集群IP" prop="clusterIP" min-width="80" />
        <el-table-column label="类型" prop="type" min-width="80" />
        <el-table-column label="创建时间" width="200">
          <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="240">
          <template #default="scope">
            <el-button icon="view" type="text" size="small" @click="viewServices(scope.row)">查看</el-button>
            <el-button icon="delete" type="text" size="small" @click="deleteServicesFunc(scope.row)">删除</el-button>
          </template>
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

      <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%" :destroy-on-close="true">
        <!-- eslint-disable-next-line vue/attribute-hyphenation -->
        <vue-code-mirror v-model:modelValue="servicesFormat" :readOnly="true" />
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { formatDate } from '@/utils/format'
import { getNamespaceOnlyName } from '@/api/kubernetes/namespace'
import { getServicesList, getServicesRaw, deleteServices } from '@/api/kubernetes/services'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'ServicesList',
  components: {
    VueCodeMirror,
  },
  setup() {
    // 响应式数据
    const searchInfo = reactive({
      namespace: ''
    })

    // 加载namespace数据
    const namespace = ref([])

    const getNamespace = async() => {
      const table = await getNamespaceOnlyName()
      if (table.code === 0) {
        namespace.value = table.data
      }
    }
    getNamespace()

    // 加载pod数据
    const tableData = ref([])
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)

    const getTableData = async() => {
      const table = await getServicesList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
      if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getTableData()

    // 查看编排
    const servicesFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewServices = async(row) => {
      const result = await getServicesRaw({ service: row.name, namespace: row.namespace })
      if (result.code === 0) {
        servicesFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getTableData()
    }

    const handleCurrentChange = (val) => {
      page.value = val
      getTableData()
    }

    // 查询
    const onSubmit = () => {
      page.value = 1
      pageSize.value = 10
      getTableData()
    }

    // 查询重置
    const onReset = () => {
      searchInfo.namespace = ''
    }

    // 删除service
    const deleteServicesFunc = async(row) => {
      ElMessageBox.confirm('此操作将永久删除该Service, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteServices({ namespace: row.namespace, service: row.name })
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '删除成功!'
            })
            const index = tableData.value.indexOf(row)
            tableData.value.splice(index, 1)
          }
        })
    }

    return {
      namespace,
      searchInfo,
      // time format
      formatDate,
      // 表格相关
      page,
      pageSize,
      total,
      tableData,
      handleSizeChange,
      handleCurrentChange,
      // 查询
      onSubmit,
      // 查询重置
      onReset,
      // 查看编排
      servicesFormat,
      dialogFormVisible,
      viewServices,
      // 删除services
      deleteServicesFunc
    }
  }
}
</script>
