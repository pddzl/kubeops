<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true">
        <el-form-item label="命名空间">
          <el-select v-model="searchInfo.namespace" clearable placeholder="请选择">
            <el-option v-for="item in namespace" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="Pod名称">
          <el-input placeholder="Pod名称" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column align="left" label="名称" min-width="220">
          <template #default="scope">
            <router-link :to="{ name: 'pod_detail', query: { name: scope.row.name } }">
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column>
          <template #default="scope" label="状态" min-width="100">
            <el-tag :type="statusTypeFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="节点" prop="node" min-width="100" />
        <el-table-column prop="resource.cpuLimit" label="CPU预留" width="100" />
        <el-table-column prop="resource.memoryLimit" label="CPU限制" width="100" />
        <el-table-column prop="resource.cpuRequests" label="内存预留" width="100" />
        <el-table-column prop="resource.memoryRequests" label="内存限制" width="100" />
        <el-table-column label="创建时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.createTimestamp) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="240">
          <template #default>
            <el-button icon="edit" type="text" size="small">编辑</el-button>
            <el-button icon="tickets" type="text" size="small">日志</el-button>
            <el-button icon="ArrowRight" type="text" size="small">终端</el-button>
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
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { formatDate } from '@/utils/format'
import { getNamespaceOnlyName } from '@/api/kubernetes/namespace'
import { getPodList } from '@/api/kubernetes/pod'
export default {
  name: 'Pod',
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
      const table = await getPodList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
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
      'Running': 'success'
    }

    const statusTypeFilter = (status) => {
      return statusTypeMap[status] || 'info'
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
      namespace,
      searchInfo,
      tableData,
      // filter
      statusTypeFilter,
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
      onReset
    }
  }
}
</script>
