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
              :to="{ name: 'pod_detail', query: { pod: scope.row.name, namespace: scope.row.namespace } }"
            >
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="命名空间" prop="namespace" min-width="120" />
        <el-table-column label="状态" min-width="100">
          <template #default="scope">
            <el-tag :type="statusPodFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="节点" prop="node" min-width="100" />
        <el-table-column label="创建时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="240">
          <template #default="scope">
            <el-button icon="view" type="primary" link size="small" @click="viewPod(scope.row)">查看</el-button>
            <el-button
              icon="tickets"
              type="primary"
              link
              size="small"
              @click="routerPod(scope.row, 'log')"
            >日志</el-button>
            <el-button icon="ArrowRight" type="primary" link size="small" @click="routerPod(scope.row, 'terminal')">终端</el-button>
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
        <vue-code-mirror v-model:modelValue="podFormat" :readOnly="true" />
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { formatDate } from '@/utils/format'
import { statusPodFilter } from '@/mixin/filter.js'
import { getNamespaceOnlyName } from '@/api/kubernetes/namespace'
import { getPodList, getPodRaw, deletePod } from '@/api/kubernetes/pod'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'Pod',
  components: {
    VueCodeMirror,
  },
  setup() {
    // search
    const searchInfo = reactive({
      namespace: ''
    })

    // 路由
    const router = useRouter()

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
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const tableData = ref([])

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

    // 编辑
    const podFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewPod = async(row) => {
      const result = await getPodRaw({ pod: row.name, namespace: row.namespace })
      if (result.code === 0) {
        podFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 跳转日志/webssh页面
    const routerPod = async(row, dest) => {
      if (dest === 'log') {
        router.push({ name: 'pod_log', query: { pod: row.name, namespace: row.namespace }})
      } else if (dest === 'terminal') {
        router.push({ name: 'pod_terminal', query: { pod: row.name, namespace: row.namespace }})
      }
    }

    // 删除
    const deleteFunc = async(row) => {
      ElMessageBox.confirm('此操作将永久删除该Pod, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deletePod({ namespace: row.namespace, pod: row.name })
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
      // data、查询相关
      searchInfo,
      namespace,
      tableData,
      onSubmit,
      onReset,
      // filter
      statusPodFilter,
      // time format
      formatDate,
      // 编辑
      podFormat,
      dialogFormVisible,
      viewPod,
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange,
      // 路由
      routerPod,
      // 删除
      deleteFunc
    }
  }
}
</script>
