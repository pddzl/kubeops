<template>
  <div>
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column label="名称">
          <template #default="scope">
            <router-link
              :to="{ name: 'clusterRoleBinding_detail', query: { clusterRoleBinding: scope.row.name } }"
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
            <el-button icon="view" type="primary" link size="small" @click="viewClusterRoleBinding(scope.row)">查看</el-button>
            <el-button icon="delete" type="primary" link size="small" @click="deleteFunc(scope.row)">删除</el-button>
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
        <vue-code-mirror v-model:modelValue="clusterRoleBindingFormat" :readOnly="true" />
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { formatDate } from '@/utils/format'
import { getClusterRoleBindingList, getClusterRoleBindingRaw, deleteClusterRoleBinding } from '@/api/kubernetes/clusterRoleBinding'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'ClusterRoleBindingList',
  components: {
    VueCodeMirror,
  },
  setup() {
    // 加载clusterRoleBinding数据
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const tableData = ref([])

    const getTableData = async() => {
      const table = await getClusterRoleBindingList({ page: page.value, pageSize: pageSize.value })
      if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getTableData()

    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getTableData()
    }

    const handleCurrentChange = (val) => {
      page.value = val
      getTableData()
    }

    // 编辑
    const clusterRoleBindingFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewClusterRoleBinding = async(row) => {
      const result = await getClusterRoleBindingRaw({ cluster_role_binding: row.name })
      if (result.code === 0) {
        clusterRoleBindingFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async(row) => {
      ElMessageBox.confirm('此操作将永久删除该ClusterRoleBinding, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteClusterRoleBinding({ cluster_role_binding: row.name })
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
      // data
      tableData,
      // time format
      formatDate,
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange,
      // 编辑
      clusterRoleBindingFormat,
      dialogFormVisible,
      viewClusterRoleBinding,
      // 删除
      deleteFunc
    }
  }
}
</script>
