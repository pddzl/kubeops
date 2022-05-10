<template>
  <div>
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column label="名称">
          <template #default="scope">
            <router-link
              :to="{ name: 'clusterRole_detail', query: { clusterRole: scope.row.name } }"
            >
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作">
          <template #default="scope">
            <el-button icon="view" type="text" size="small" @click="viewClusterRole(scope.row)">查看</el-button>
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
        <vue-code-mirror v-model:modelValue="clusterRoleFormat" :readOnly="true" />
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { formatDate } from '@/utils/format'
import { getNamespaceOnlyName } from '@/api/kubernetes/namespace'
import { getClusterRoleList, getClusterRoleRaw } from '@/api/kubernetes/clusterRole'
import VueCodeMirror from '@/components/codeMirror/index.vue'
export default {
  name: 'ClusterRoleList',
  components: {
    VueCodeMirror,
  },
  setup() {
    // 响应式数据
    const namespace = ref([])
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const tableData = ref([])
    const clusterRoleFormat = ref({})
    const dialogFormVisible = ref(false)

    // 加载namespace数据
    const getNamespace = async() => {
      const table = await getNamespaceOnlyName()
      if (table.code === 0) {
        namespace.value = table.data
      }
    }
    getNamespace()

    // 加载clusterRole数据
    const getTableData = async() => {
      const table = await getClusterRoleList({ page: page.value, pageSize: pageSize.value })
      if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getTableData()

    // 操作
    const viewClusterRole = async(row) => {
      const result = await getClusterRoleRaw({ clusterRole: row.name, namespace: row.namespace })
      if (result.code === 0) {
        clusterRoleFormat.value = JSON.stringify(result.data)
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

    return {
      // 响应式数据
      namespace,
      tableData,
      clusterRoleFormat,
      dialogFormVisible,
      // time format
      formatDate,
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange,
      viewClusterRole
    }
  }
}
</script>
