<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewClusterRoleBinding">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="clusterRoleBindingDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="clusterRoleBindingDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="clusterRoleBindingDetail.roleRef && JSON.stringify(clusterRoleBindingDetail.roleRef) !== '{}'" title="Role" name="role">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>Role Reference</p>
                <router-link :to="{ name: 'clusterRole_detail', query: { clusterRole: clusterRoleBindingDetail.roleRef.name } }">
                  <el-link type="primary" :underline="false">{{ clusterRoleBindingDetail.roleRef.name }}</el-link>
                </router-link>
              </div>
              <div class="item">
                <p>apiGroup</p>
                <span>{{ clusterRoleBindingDetail.roleRef.apiGroup }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="clusterRoleBindingDetail.subjects && clusterRoleBindingDetail.subjects.length > 0" title="Subjects" name="subjects">
          <div class="info-table">
            <el-table :data="clusterRoleBindingDetail.subjects">
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
      <vue-code-mirror v-model:modelValue="clusterRoleBindingFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getClusterRoleBindingDetail, getClusterRoleBindingRaw, deleteClusterRoleBinding } from '@/api/kubernetes/clusterRoleBinding'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'ClusterRoleBindingDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'role', 'subjects'])

    // 路由
    const route = useRoute()
    const clusterRoleBinding = route.query.clusterRoleBinding

    // 加载clusterRoleBinding详情
    const clusterRoleBindingDetail = ref({})

    const getData = async() => {
      await getClusterRoleBindingDetail({ cluster_role_binding: clusterRoleBinding }).then(response => {
        if (response.code === 0) {
          clusterRoleBindingDetail.value = response.data
        }
      })
    }
    getData()

    // 编辑
    const clusterRoleBindingFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewClusterRoleBinding = async() => {
      const result = await getClusterRoleBindingRaw({ cluster_role_binding: clusterRoleBinding })
      if (result.code === 0) {
        clusterRoleBindingFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该ClusterRoleBinding, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteClusterRoleBinding({ cluster_role_binding: clusterRoleBinding })
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
      clusterRoleBindingDetail,
      // time format
      formatDate,
      // 编辑
      viewClusterRoleBinding,
      clusterRoleBindingFormat,
      dialogFormVisible,
      // 删除
      deleteFunc
    }
  }
}
</script>
