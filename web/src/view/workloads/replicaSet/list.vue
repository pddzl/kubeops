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
        <el-table-column label="名称" min-width="220">
          <template #default="scope">
            <router-link
              :to="{ name: 'replicaSet_detail', query: { replicaSet: scope.row.name, namespace: scope.row.namespace } }"
            >
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template>
        </el-table-column>
        <el-table-column label="命名空间" prop="namespace" min-width="120" />
        <el-table-column label="Pods" prop="pods" min-width="80">
          <template #default="scope">{{ scope.row.availableReplicas }} / {{ scope.row.replicas }}</template>
        </el-table-column>
        <el-table-column label="创建时间" width="200">
          <template #default="scope">{{ formatDate(scope.row.creationTimestamp) }}</template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="240">
          <template #default="scope">
            <el-button icon="view" type="primary" link size="small" @click="viewReplicaSet(scope.row)">查看</el-button>
            <el-button icon="expand" type="primary" link size="small" @click="openScaleDialog(scope.row)">伸缩</el-button>
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
      <!-- 查看对话框 -->
      <el-dialog v-model="dialogViewVisible" title="查看资源" width="55%" :destroy-on-close="true">
        <!-- eslint-disable-next-line vue/attribute-hyphenation -->
        <vue-code-mirror v-model:modelValue="replicaSetFormat" :readOnly="true" />
      </el-dialog>
      <!-- 伸缩对话框 -->
      <el-dialog v-model="dialogScaleVisible" title="伸缩" width="55%" center>
        <p style="font-weight: bold;">ReplicaSet {{ replicasetName }} will be updated to reflect the
          desired replicas count.</p>
        <div style="margin: 25px 0 25px 0px;">
          <span style="margin-right: 10px;">Desired Replicas:</span>
          <el-input-number v-model="desiredNum" :min="0" :max="50" style="margin-right: 20px;" />
          <span style="margin-right: 10px;">Actual Replicas: </span>
          <el-input-number v-model="ActualNum" disabled />
        </div>
        <warning-bar :title="warningTitle" />
        <template #footer>
          <el-button @click="closeScaleDialog">取消</el-button>
          <el-button type="primary" @click="scaleFunc">确认</el-button>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { ref, reactive, watch } from 'vue'
import { formatDate } from '@/utils/format'
import { getNamespaceOnlyName } from '@/api/kubernetes/namespace'
import { getReplicaSetList, getReplicaSetRaw, deleteReplicaSet } from '@/api/kubernetes/replicaSet'
import { scale } from '@/api/kubernetes/scale'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import warningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'ReplicaSetList',
  components: {
    VueCodeMirror,
    warningBar
  },
  setup() {
    // 加载namespace by name
    const namespace = ref([])

    const getNamespace = async() => {
      const table = await getNamespaceOnlyName()
      if (table.code === 0) {
        namespace.value = table.data
      }
    }
    getNamespace()

    // 加载pod数据
    const searchInfo = reactive({
      namespace: ''
    })
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const tableData = ref([])

    const getTableData = async() => {
      const table = await getReplicaSetList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
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

    // 查看编排
    const dialogViewVisible = ref(false)
    const replicaSetFormat = ref({})

    const viewReplicaSet = async(row) => {
      const result = await getReplicaSetRaw({ replicaSet: row.name, namespace: row.namespace })
      if (result.code === 0) {
        replicaSetFormat.value = JSON.stringify(result.data)
      }
      dialogViewVisible.value = true
    }

    // 伸缩
    const dialogScaleVisible = ref(false)
    const warningTitle = ref('')
    const replicasetName = ref('')
    const desiredNum = ref(0)
    const ActualNum = ref(0)
    const activeRow = ref({})

    // -> 打开对话框
    const openScaleDialog = (row) => {
      replicasetName.value = row.name
      desiredNum.value = row.replicas
      ActualNum.value = row.replicas
      activeRow.value = row
      dialogScaleVisible.value = true
      warningTitle.value = `This action is equivalent to: kubectl scale -n ${row.namespace} replicaset ${row.name} --replicas=${row.replicas}`
    }

    watch(desiredNum, (val) => {
      warningTitle.value = `This action is equivalent to: kubectl scale -n ${activeRow.value.namespace} replicaset ${activeRow.value.name} --replicas=${val}`
    })

    // -> 关闭对话框
    const closeScaleDialog = () => {
      dialogScaleVisible.value = false
    }

    // -> 操作
    const scaleFunc = async() => {
      const res = await scale({ namespace: activeRow.value.namespace, name: activeRow.value.name, kind: 'replicaSet', num: desiredNum.value })
      if (res.code === 0) {
        const index = tableData.value.indexOf(activeRow.value)
        tableData.value[index].availableReplicas = desiredNum.value
        tableData.value[index].replicas = desiredNum.value
        ElMessage({
          type: 'success',
          message: '伸缩成功',
          showClose: true
        })
      }
      dialogScaleVisible.value = false
    }

    // 删除
    const deleteFunc = async(row) => {
      ElMessageBox.confirm('此操作将永久删除该ReplicaSet, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteReplicaSet({ namespace: row.namespace, replicaSet: row.name })
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
      // 表单数据相关
      namespace,
      searchInfo,
      tableData,
      // 查看编排
      replicaSetFormat,
      dialogViewVisible,
      viewReplicaSet,
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
      // 伸缩
      dialogScaleVisible,
      warningTitle,
      replicasetName,
      desiredNum,
      ActualNum,
      activeRow,
      openScaleDialog,
      closeScaleDialog,
      scaleFunc,
      // 删除
      deleteFunc
    }
  }
}
</script>
