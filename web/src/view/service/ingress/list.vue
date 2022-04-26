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
        <el-table-column label="名称" min-width="120">
          <template #default="scope">
            <router-link
              :to="{ name: 'services_detail', query: { service: scope.row.name, namespace: scope.row.namespace } }"
            >
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="命名空间" prop="namespace" min-width="120" />
        <el-table-column label="EndPoints" min-width="200">
          <template #default="scope">
            <el-tag v-for="dp in scope.row.endPoints" :key="dp" style="margin-right: 5px;">{{ dp }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Hosts" min-width="200">
          <template #default="scope">
            <el-tag v-for="host in scope.row.hosts" :key="host" style="margin-right: 5px;">{{ host }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="200">
          <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="240">
          <template #default="scope">
            <el-button icon="view" type="text" size="small" @click="viewIngress(scope.row)">查看</el-button>
            <el-button icon="delete" type="text" size="small">删除</el-button>
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

      <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
        <!-- eslint-disable-next-line vue/attribute-hyphenation -->
        <vue-code-mirror v-model:modelValue="ingressFormat" :readOnly="true" />
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { formatDate } from '@/utils/format'
import { getNamespaceOnlyName } from '@/api/kubernetes/namespace'
import { getIngressList, getIngressRaw } from '@/api/kubernetes/ingress'
import VueCodeMirror from '@/components/codeMirror/index.vue'
export default {
  name: 'IngressList',
  components: {
    VueCodeMirror,
  },
  setup() {
    // 响应式数据
    const namespace = ref([])
    const searchInfo = reactive({
      namespace: ''
    })
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const tableData = ref([])
    const ingressFormat = ref({})
    const dialogFormVisible = ref(false)

    // 加载namespace数据
    const getNamespace = async() => {
      const table = await getNamespaceOnlyName()
      if (table.code === 0) {
        namespace.value = table.data
      }
    }
    getNamespace()

    // 加载pod数据
    const getTableData = async() => {
      const table = await getIngressList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
      if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getTableData()

    // 操作
    const viewIngress = async(row) => {
      const result = await getIngressRaw({ ingress: row.name, namespace: row.namespace })
      if (result.code === 0) {
        ingressFormat.value = JSON.stringify(result.data)
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

    // 重置
    const onReset = () => {
      searchInfo.namespace = ''
    }

    return {
      // 响应式数据
      namespace,
      searchInfo,
      tableData,
      ingressFormat,
      dialogFormVisible,
      // time format
      formatDate,
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange,
      // 查询
      onSubmit,
      onReset,
      viewIngress
    }
  }
}
</script>
