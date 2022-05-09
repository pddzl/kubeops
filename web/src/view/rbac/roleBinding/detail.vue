<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewRoleBinding">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="roleBindingDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="roleBindingDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="roleBindingDetail.roleRef && JSON.stringify(roleBindingDetail.roleRef) !== '{}'" title="Role" name="role">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>Role Reference</p>
                <router-link :to="{ name: 'role_detail', query: { role: roleBindingDetail.roleRef.name, namespace: namespace } }">
                  <el-link type="primary" :underline="false">{{ roleBindingDetail.roleRef.name }}</el-link>
                </router-link>
              </div>
              <div class="item">
                <p>apiGroup</p>
                <span>{{ roleBindingDetail.roleRef.apiGroup }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="roleBindingFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getRoleBindingDetail, getRoleBindingRaw } from '@/api/kubernetes/roleBinding'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
export default {
  name: 'RoleBindingDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    const activeNames = ref(['metadata', 'role'])
    const roleBindingDetail = ref({})
    const roleBindingFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const roleBinding = route.query.roleBinding

    // 加载roleBinding详情
    const getData = async() => {
      await getRoleBindingDetail({ namespace: namespace, roleBinding: roleBinding }).then(response => {
        if (response.code === 0) {
          roleBindingDetail.value = response.data
        }
      })
    }
    getData()

    // 操作
    const viewRoleBinding = async() => {
      const result = await getRoleBindingRaw({ roleBinding: roleBinding, namespace: namespace })
      if (result.code === 0) {
        roleBindingFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      roleBindingDetail,
      formatDate,
      dialogFormVisible,
      namespace,
      // 操作
      viewRoleBinding,
      roleBindingFormat
    }
  }
}
</script>
