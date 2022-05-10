<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewClusterRole">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="clusterRoleDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="clusterRoleDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="clusterRoleDetail.rules && clusterRoleDetail.rules.length > 0" title="规则" name="rules">
          <div class="info-table">
            <el-table :data="clusterRoleDetail.rules">
              <el-table-column label="资源" min-width="200px">
                <template #default="scope">
                  <span v-for="(rs,index) in scope.row.resources" :key="index">
                    {{ rs }}<span v-if="index !== scope.row.resources.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
              <el-table-column label="非资源URL">
                <template #default="scope">
                  <span v-for="(nRsUrl,index) in scope.row.nonResourceURLs" :key="index">
                    {{ nRsUrl }}<span v-if="index !== scope.row.nonResourceURLs.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
              <el-table-column label="资源名">
                <template #default="scope">
                  <span v-for="(rn,index) in scope.row.resourceNames" :key="index">
                    {{ rn }}<span v-if="index !== scope.row.resourceNames.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
              <el-table-column label="动作">
                <template #default="scope">
                  <span v-for="(v,index) in scope.row.verbs" :key="index">
                    {{ v }}<span v-if="index !== scope.row.verbs.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
              <el-table-column label="API组">
                <template #default="scope">
                  <span v-for="(v,index) in scope.row.apiGroups" :key="index">
                    {{ v }}<span v-if="index !== scope.row.apiGroups.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="clusterRoleFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getClusterRoleDetail, getClusterRoleRaw } from '@/api/kubernetes/clusterRole'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
export default {
  name: 'ClusterRoleDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    const activeNames = ref(['metadata', 'rules'])
    const clusterRoleDetail = ref({})
    const clusterRoleFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const clusterRole = route.query.clusterRole

    // 加载clusterRole详情
    const getData = async() => {
      await getClusterRoleDetail({ clusterRole: clusterRole }).then(response => {
        if (response.code === 0) {
          clusterRoleDetail.value = response.data
        }
      })
    }
    getData()

    // 操作
    const viewClusterRole = async() => {
      const result = await getClusterRoleRaw({ clusterRole: clusterRole })
      if (result.code === 0) {
        clusterRoleFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      clusterRoleDetail,
      formatDate,
      dialogFormVisible,
      // 操作
      viewClusterRole,
      clusterRoleFormat
    }
  }
}
</script>
