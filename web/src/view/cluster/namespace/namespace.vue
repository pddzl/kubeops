<template>
  <div>
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column align="left" label="名称" min-width="180" prop="name" />
        <el-table-column align="left" label="状态" width="100" prop="status">
          <template #default="scope">
            <el-tag :type="statusTypeFilter(scope.row.status)" size="small">
              {{ scope.row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="标签" min-width="400">
          <template #default="scope">
            <span v-for="(value, key) in scope.row.labels" :key="key" class="span-shadow">
              {{ key }}: {{ value }}
            </span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" min-width="200">
          <template #default="scope">{{ scope.row.createTimestamp }}</template>
        </el-table-column>
        <el-table-column align="left" fixed="right" label="操作" min-width="160">
          <template #default="scope">
            <el-button
              icon="edit"
              size="small"
              type="text"
              @click="editNode(scope.row)"
            >编辑</el-button>
            <el-button
              icon="delete"
              size="small"
              type="text"
              @click="deleteNode(scope.row)"
            >删除</el-button>
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
  </div>
</template>

<script>
import { ref } from 'vue'
import { getNamespaceList } from '@/api/kubernetes/namespace'
import { formatDate } from '@/utils/format'
// import '@/style/kop.scss'
export default {
  name: 'Namespace',
  setup() {
    // 响应式数据
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const tableData = ref([])

    // 加载namespace数据
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

    // filter
    const statusTypeMap = {
      'Active': 'success',
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

    const statusTypeFilter = (status) => {
      return statusTypeMap[status] || 'info'
    }

    return {
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange,
      // 日期处理
      formatDate,
      // 数据
      tableData,
      // filter
      statusTypeFilter
    }
  }
}
</script>

<style scoped lang="scss">
</style>
