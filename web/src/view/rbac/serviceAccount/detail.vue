<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewServiceAccount">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
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
import { getServiceAccountDetail, getServiceAccountRaw } from '@/api/kubernetes/serviceAccount'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
export default {
  name: 'ServiceAccountDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    const activeNames = ref(['metadata', 'secrets'])
    const serviceAccountDetail = ref({})
    const serviceAccountFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const serviceAccount = route.query.serviceAccount

    // 加载serviceAccount详情
    const getData = async() => {
      await getServiceAccountDetail({ namespace: namespace, serviceAccount: serviceAccount }).then(response => {
        if (response.code === 0) {
          serviceAccountDetail.value = response.data
          console.log('serviceAccountDetail', serviceAccountDetail.value)
        }
      })
    }
    getData()

    // 操作
    const viewServiceAccount = async() => {
      const result = await getServiceAccountRaw({ serviceAccount: serviceAccount, namespace: namespace })
      if (result.code === 0) {
        serviceAccountFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      serviceAccountDetail,
      formatDate,
      dialogFormVisible,
      namespace,
      // 操作
      viewServiceAccount,
      serviceAccountFormat
    }
  }
}
</script>
