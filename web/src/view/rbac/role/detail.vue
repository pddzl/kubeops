<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewRole">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="roleDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="roleDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="roleDetail.rules && roleDetail.rules.length > 0" title="规则" name="rules">
          <div class="info-table">
            <el-table :data="roleDetail.rules">
              <el-table-column label="资源">
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
      <vue-code-mirror v-model:modelValue="roleFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getRoleDetail, getRoleRaw } from '@/api/kubernetes/role'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
export default {
  name: 'RoleDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    const activeNames = ref(['metadata', 'rules'])
    const roleDetail = ref({})
    const roleFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const role = route.query.role

    // 加载role详情
    const getData = async() => {
      await getRoleDetail({ namespace: namespace, role: role }).then(response => {
        if (response.code === 0) {
          roleDetail.value = response.data
          console.log('roleDetail', roleDetail.value)
        }
      })
    }
    getData()

    // 操作
    const viewRole = async() => {
      const result = await getRoleRaw({ role: role, namespace: namespace })
      if (result.code === 0) {
        roleFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      roleDetail,
      formatDate,
      dialogFormVisible,
      namespace,
      // 操作
      viewRole,
      roleFormat
    }
  }
}
</script>
