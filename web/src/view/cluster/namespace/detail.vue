<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewNamespace">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="namespaceDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="namespaceDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>状态</p>
                <span class="content">
                  <el-tag :type="statusNsFilter(namespaceDetail.status)" size="small">{{ namespaceDetail.status }}
                  </el-tag>
                </span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item
          v-if="namespaceDetail.resourceQuotaList && namespaceDetail.resourceQuotaList.length > 0"
          title="资源配额"
          name="resourceQuota"
        >
          <div style="margin-right: 20px;">
            <el-table :data="namespaceDetail.resourceQuotaList">
              <el-table-column label="名称" align="center">
                <template #default="scope">
                  <el-tag size="small" effect="plain">{{ scope.row.objectMeta.name }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="limits" align="center">
                <el-table-column label="cpu used" align="center" prop="statusList.limits_cpu.used" />
                <el-table-column label="cpu hard" align="center" prop="statusList.limits_cpu.hard" />
                <el-table-column label="memory used" align="center" prop="statusList.limits_memory.used" />
                <el-table-column label="memory hard" align="center" prop="statusList.limits_memory.hard" />
              </el-table-column>
              <el-table-column label="requests" align="center">
                <el-table-column label="cpu used" align="center" prop="statusList.requests_cpu.used" />
                <el-table-column label="cpu hard" align="center" prop="statusList.requests_cpu.hard" />
                <el-table-column label="memory used" align="center" prop="statusList.requests_memory.used" />
                <el-table-column label="memory hard" align="center" prop="statusList.requests_memory.hard" />
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item
          v-if="namespaceDetail.resourceLimits && namespaceDetail.resourceLimits.length > 0"
          title="资源限制"
          name="resourceLimit"
        >
          <div v-for="(rLimit, keyRl) in namespaceDetail.resourceLimits" :key="keyRl" style="margin-right: 20px;">
            <div style="margin-bottom: 10px;">
              <el-tag size="small">{{ rLimit.name }}</el-tag>
            </div>
            <el-table :data="rLimit.limits">
              <el-table-column label="类型" align="center" prop="type" />
              <el-table-column label="max" align="center">
                <el-table-column label="cpu" align="center" prop="max.cpu" />
                <el-table-column label="memory" align="center" prop="max.memory" />
              </el-table-column>
              <el-table-column label="min" align="center">
                <el-table-column label="cpu" align="center" prop="min.cpu" />
                <el-table-column label="memory" align="center" prop="min.memory" />
              </el-table-column>
              <el-table-column label="default" align="center">
                <el-table-column label="cpu" align="center" prop="default.cpu" />
                <el-table-column label="memory" align="center" prop="default.memory" />
              </el-table-column>
              <el-table-column label="defaultRequest" align="center">
                <el-table-column label="cpu" align="center" prop="defaultRequest.cpu" />
                <el-table-column label="memory" align="center" prop="defaultRequest.memory" />
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="namespaceFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { formatDate } from '@/utils/format'
import { statusNsFilter } from '@/mixin/filter'
import { getNamespaceDetail, getNamespaceRaw, deleteNamespace } from '@/api/kubernetes/namespace'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'NamespaceDetail',
  components: {
    MetaData,
    VueCodeMirror
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'spec', 'resourceQuota', 'resourceLimit'])

    // 路由
    const route = useRoute()
    const namespace = route.query.name

    // 加载namespace详情
    const namespaceDetail = ref({})

    const getData = async() => {
      await getNamespaceDetail({ name: namespace }).then(response => {
        if (response.code === 0) {
          namespaceDetail.value = response.data
        }
      })
    }
    getData()

    // 编辑
    const namespaceFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewNamespace = async() => {
      const res = await getNamespaceRaw({ name: namespace })
      if (res.code === 0) {
        namespaceFormat.value = JSON.stringify(res.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该Namespace, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteNamespace({ name: namespace })
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '删除成功!'
            })
          }
        })
    }

    return {
      // 折叠面板
      activeNames,
      // data
      namespace,
      namespaceDetail,
      // filter
      statusNsFilter,
      // time format
      formatDate,
      // 编辑
      dialogFormVisible,
      namespaceFormat,
      viewNamespace,
      // 删除
      deleteFunc
    }
  }
}
</script>
