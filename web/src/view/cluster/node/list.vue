<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus">扩容</el-button>
        <el-button
          icon="delete"
          size="small"
          :disabled="!nodes.length"
          style="margin-left: 10px;"
        >删除</el-button>
      </div>
      <el-table
        :data="tableData"
        @sort-change="sortChange"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="名称" min-width="100">
          <template #default="scope">
            <router-link :to="{ name: 'node_detail', query: { name: scope.row.name } }">
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="内部IP" min-width="130" prop="internalIP" />
        <el-table-column label="角色" min-width="220" prop="roles">
          <template #default="scope">
            <span v-for="(role, index) in scope.row.role" :key="index" style="margin-right: 5px;">
              <el-tag size="small">{{ role }}</el-tag>
            </span>
          </template>
        </el-table-column>
        <el-table-column label="状态" min-width="120" prop="status">
          <template #default="scope">
            <el-tag
              :type="statusNodeTypeFilter(scope.row.status)"
              size="small"
            >{{ statusNodeFilter(scope.row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="CPU (core)" min-width="100" prop="cpu" />
        <el-table-column label="内存 (GB)" min-width="100">
          <template
            #default="scope"
          >{{ (parseInt(scope.row.memory.slice(0, -2)) / 1024 / 1024).toFixed(2) }}</template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          min-width="180"
          prop="metadata.creationTimestamp"
          sortable="custom"
        >
          <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="200">
          <template #default="scope">
            <el-button icon="view" size="small" type="primary" link @click="viewNode(scope.row)">查看</el-button>
            <el-button icon="delete" size="small" type="primary" link @click="deleteNode(scope.row)">删除</el-button>
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

    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%" :destroy-on-close="true">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="nodeFormat" :readOnly="true" />
      <!-- <warning-bar title="此操作相当于：kubectl apply -f <spec.json>" style="margin-top: 10px;" /> -->
    </el-dialog>
  </div>
</template>

<script>
import 'vue-json-pretty/lib/styles.css'
import { getNodeList, getNodeRaw } from '@/api/kubernetes/node'
import { toSQLLine } from '@/utils/stringFun'
// import warningBar from '@/components/warningBar/warningBar.vue'
import { formatDate } from '@/utils/format'
import { ref, reactive } from 'vue'
import { ElMessageBox } from 'element-plus'
import { statusNodeTypeFilter, statusNodeFilter } from '@/mixin/filter.js'
import VueCodeMirror from '@/components/codeMirror/index.vue'
export default {
  name: 'Node',
  components: {
    VueCodeMirror,
    // warningBar
  },
  setup() {
    // 响应式数据
    const page = ref(1)
    const total = ref(0)
    const pageSize = ref(10)
    const tableData = ref([])
    const searchInfo = reactive({})
    const dialogFormVisible = ref(false)
    const nodeFormat = ref({})
    const state = reactive({
      virtualLines: 30,
      showLength: true
    })
    const nodes = ref([])

    // 加载node数据
    const getTableData = async() => {
      const table = await getNodeList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
      if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }

    getTableData()

    // 操作
    const viewNode = async(row) => {
      const result = await getNodeRaw({ name: row.name })
      if (result.code === 0) {
        nodeFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    const deleteNode = async(row) => {
      ElMessageBox.confirm(`此操作将永久删除节点${row.metadata.name}, 是否继续?`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          getTableData()
        })
    }

    // 批量操作
    const handleSelectionChange = (val) => {
      nodes.value = val
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

    // 排序
    const sortChange = ({ prop, order }) => {
      if (prop) {
        searchInfo.orderKey = toSQLLine(prop)
        searchInfo.desc = order === 'descending'
      }
      getTableData()
    }

    return {
      // 响应式数据
      page,
      pageSize,
      total,
      tableData,
      dialogFormVisible,
      state,
      nodes,
      nodeFormat,
      // 操作
      viewNode,
      deleteNode,
      handleSelectionChange,
      // 分页
      handleCurrentChange,
      handleSizeChange,
      // 排序
      sortChange,
      // 过滤器
      statusNodeTypeFilter,
      statusNodeFilter,
      // 时间格式化
      formatDate
    }
  }
}
</script>
