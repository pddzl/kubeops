<template>
  <div>
    <div class="detail-operation">
      <el-button icon="view" size="small" type="primary" plain @click="viewServiceAccount">查看</el-button>
      <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="serviceAccountDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="serviceAccountDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="serviceAccountDetail.secrets && serviceAccountDetail.secrets.length > 0" title="Secrets" name="secrets">
          <p v-for="secret in serviceAccountDetail.secrets" :key="secret">
            <router-link :to="{ name: 'secret_detail', query: { secret: secret.name, namespace: namespace } }">
              <el-link type="primary" :underline="false">{{ secret.name }}</el-link>
            </router-link>
          </p>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="serviceAccountFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getServiceAccountDetail, getServiceAccountRaw, deleteServiceAccount } from '@/api/kubernetes/serviceAccount'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'ServiceAccountDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'secrets'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const serviceAccount = route.query.serviceAccount

    // 加载serviceAccount详情
    const serviceAccountDetail = ref({})

    const getData = async() => {
      await getServiceAccountDetail({ namespace: namespace, serviceAccount: serviceAccount }).then(response => {
        if (response.code === 0) {
          serviceAccountDetail.value = response.data
          console.log('serviceAccountDetail', serviceAccountDetail.value)
        }
      })
    }
    getData()

    // 编辑
    const serviceAccountFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewServiceAccount = async() => {
      const result = await getServiceAccountRaw({ serviceAccount: serviceAccount, namespace: namespace })
      if (result.code === 0) {
        serviceAccountFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该ServiceAccount, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteServiceAccount({ namespace: namespace, serviceAccount: serviceAccount })
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
      serviceAccountDetail,
      // time format
      formatDate,
      // 操作
      dialogFormVisible,
      serviceAccountFormat,
      viewServiceAccount,
      // 删除
      deleteFunc
    }
  }
}
</script>
