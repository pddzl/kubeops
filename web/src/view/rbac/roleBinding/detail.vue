<template>
  <div>
    <div class="detail-operation">
      <el-button icon="view" size="small" type="primary" plain @click="viewRoleBinding">查看</el-button>
      <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
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
        <el-collapse-item v-if="roleBindingDetail.subjects && roleBindingDetail.subjects.length > 0" title="Subjects" name="subjects">
          <div class="info-table">
            <el-table :data="roleBindingDetail.subjects">
              <el-table-column label="名称" prop="name" />
              <el-table-column label="命名空间" prop="namespace" />
              <el-table-column label="类型" prop="kind" />
              <el-table-column label="apiGroup" prop="apiGroup" />
            </el-table>
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
import { getRoleBindingDetail, getRoleBindingRaw, deleteRoleBinding } from '@/api/kubernetes/roleBinding'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'RoleBindingDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'role', 'subjects'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const roleBinding = route.query.roleBinding

    // 加载roleBinding详情
    const roleBindingDetail = ref({})

    const getData = async() => {
      await getRoleBindingDetail({ namespace: namespace, roleBinding: roleBinding }).then(response => {
        if (response.code === 0) {
          roleBindingDetail.value = response.data
        }
      })
    }
    getData()

    // 编辑
    const roleBindingFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewRoleBinding = async() => {
      const result = await getRoleBindingRaw({ roleBinding: roleBinding, namespace: namespace })
      if (result.code === 0) {
        roleBindingFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该RoleBinding, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteRoleBinding({ namespace: namespace, roleBinding: roleBinding })
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
      roleBindingDetail,
      // time format
      formatDate,
      // 编辑
      dialogFormVisible,
      roleBindingFormat,
      viewRoleBinding,
      // 删除
      deleteFunc
    }
  }
}
</script>
