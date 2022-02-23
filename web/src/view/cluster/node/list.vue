<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus">扩容</el-button>
        <el-button icon="delete" size="mini" :disabled="!nodes.length" style="margin-left: 10px;">删除</el-button>
      </div>
      <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="名称" min-width="100">
          <template #default="scope">
            <router-link :to="{name:'node_detail', query:{name:scope.row.metadata.name}}">
              <el-link type="primary" :underline="false">{{ scope.row.metadata.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column align="left" label="内部IP" min-width="120" prop="status.addresses[0].address" />
        <el-table-column align="left" label="角色" min-width="230" prop="roles">
          <template #default="scope">
            <span v-for="(role, index) in scope.row.roles" :key="index" style="margin-right: 5px;">
              <el-tag size="small">{{ role }}</el-tag>
            </span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" min-width="150" prop="ready">
          <template #default="scope">
            <el-tag :type="statusTypeFilter(scope.row.ready)" size="small">
              {{ statusFilter(scope.row.ready) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" min-width="150" prop="metadata.creationTimestamp" sortable="custom">
          <template #default="scope">{{ formatDate(scope.row.metadata.creationTimestamp) }}</template>
        </el-table-column>

        <el-table-column align="left" fixed="right" label="操作" width="200">
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

    <el-dialog v-model="dialogFormVisible" title="编辑资源">
      <el-tabs type="border-card">
        <el-tab-pane label="JSON">
          <vue-json-pretty
            style="height:400px"
            :virtual="true"
            :virtual-lines="+state.virtualLines"
            :show-length="state.showLength"
            :data="state.data"
          />
        </el-tab-pane>
      </el-tabs>
      <warning-bar title="此操作相当于：kubectl apply -f <spec.json>" style="margin-top: 10px;" />
    </el-dialog>

  </div>
</template>

<script>
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'
import { getNodeList } from '@/api/kubernetes/node'
import { toSQLLine } from '@/utils/stringFun'
import warningBar from '@/components/warningBar/warningBar.vue'
import { formatDate } from '@/utils/format'
import { ref, reactive } from 'vue'
import { ElMessageBox } from 'element-plus'
export default {
  name: 'Node',
  components: {
    VueJsonPretty,
    warningBar
  },
  setup() {
    // 响应式数据
    const page = ref(1)
    const total = ref(0)
    const pageSize = ref(10)
    const tableData = ref([])
    const searchInfo = reactive({})
    const dialogFormVisible = ref(false)
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
      const roleRe = new RegExp('node-role.kubernetes.io\/(.*)')
      tableData.value.forEach(element => {
        // 角色
        element.roles = []
        Object.keys(element.metadata.labels).forEach(item => {
          const role = roleRe.exec(item)
          if (role) {
            element.roles.push(role[1])
          }
        })
        // 状态
        for (const item of element.status.conditions) {
          if (item.type === 'Ready') {
            element.ready = item.status
            break
          }
        }
      })
    }

    getTableData()

    // 操作
    const editNode = async(row) => {
      const nodeData = JSON.parse(JSON.stringify(row))
      delete nodeData.roles
      delete nodeData.ready
      state.val = JSON.stringify(nodeData)
      state.data = nodeData
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

    // statusFilter
    const statusTypeMap = {
      'True': 'success',
      'Unknown': 'danger',
    }

    const statusTypeFilter = (status) => {
      return statusTypeMap[status] || 'info'
    }

    const statusMap = {
      'True': 'Ready',
      'Unknown': 'NotReady'
    }

    const statusFilter = (status) => {
      return statusMap[status] || 'Unknown'
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
      // 操作
      editNode,
      deleteNode,
      handleSelectionChange,
      // 分页
      handleCurrentChange,
      handleSizeChange,
      // 排序
      sortChange,
      // 过滤器
      statusTypeFilter,
      statusFilter,
      formatDate
    }
  }
}
</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.warning {
  color: #dc143c;
}
</style>
