<template>
  <div>
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column label="名称">
          <template #default="scope">
            <router-link :to="{ name: 'namespace_detail', query: { name: scope.row.name } }">
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="状态">
          <template #default="scope">
            <el-tag :type="statusNsFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template #default="scope">{{ scope.row.creationTimestamp }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作">
          <template #default="scope">
            <el-button icon="view" size="small" type="text" @click="editNamespace(scope.row)">查看</el-button>
            <el-button icon="delete" size="small" type="text" @click="deleteFunc(scope.row)">删除</el-button>
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
    </div>

    <el-dialog v-model="dialogFormVisible" title="编辑资源" :destroy-on-close="true">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="namespaceFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { statusNsFilter } from '@/mixin/filter'
import { getNamespaceList, getNamespaceRaw, deleteNamespace } from '@/api/kubernetes/namespace'
import { formatDate } from '@/utils/format'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'Namespace',
  components: {
    VueCodeMirror
  },
  setup() {
    // 加载namespace list
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const tableData = ref([])

    const getTableData = async() => {
      const table = await getNamespaceList({ page: page.value, pageSize: pageSize.value })
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
    const dialogFormVisible = ref(false)
    const namespaceFormat = ref({})

    const editNamespace = async(row) => {
      const res = await getNamespaceRaw({ name: row.name })
      if (res.code === 0) {
        namespaceFormat.value = JSON.stringify(res.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async(row) => {
      ElMessageBox.confirm('此操作将永久删除该Namespace, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteNamespace({ name: row.name })
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
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange,
      // 日期处理
      formatDate,
      // filter
      statusNsFilter,
      // 操作
      dialogFormVisible,
      namespaceFormat,
      editNamespace,
      // 删除
      deleteFunc
    }
  }
}
</script>
