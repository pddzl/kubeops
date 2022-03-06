<template>
  <div>
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column label="名称" min-width="180">
          <template #default="scope">
            <router-link :to="{ name: 'namespace_detail', query: { name: scope.row.name } }">
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="120" prop="status">
          <template #default="scope">
            <el-tag :type="statusNsFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="标签" min-width="400">
          <template #default="scope">
            <span
              v-for="(value, key) in scope.row.labels"
              :key="key"
              class="span-shadow"
            >{{ key }}: {{ value }}</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" min-width="200">
          <template #default="scope">{{ scope.row.creationTimestamp }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" min-width="160">
          <template #default="scope">
            <el-button icon="view" size="small" type="text" @click="editNamespace(scope.row)">查看</el-button>
            <el-button icon="delete" size="small" type="text" @click="deleteNamespace(scope.row)">删除</el-button>
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

    <el-dialog v-model="dialogFormVisible" title="编辑资源">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="namespaceFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { statusNsFilter } from '@/mixin/filter'
import { getNamespaceList, getNamespaceRaw } from '@/api/kubernetes/namespace'
import { formatDate } from '@/utils/format'
import VueCodeMirror from '@/components/codeMirror/index.vue'
export default {
  name: 'Namespace',
  components: {
    VueCodeMirror
  },
  setup() {
    // 响应式数据
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const tableData = ref([])
    const dialogFormVisible = ref(false)
    const namespaceFormat = ref({})

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

    // 操作
    const editNamespace = async(row) => {
      // namespaceFormat.value = JSON.stringify(row)
      const result = await getNamespaceRaw({ name: row.name })
      if (result.code === 0) {
        namespaceFormat.value = JSON.stringify(result.data)
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
      dialogFormVisible,
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
      statusNsFilter,
      // 操作
      editNamespace,
      namespaceFormat
    }
  }
}
</script>

<style scoped lang="scss">
</style>
