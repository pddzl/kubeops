<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewConfigMap">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="configMapDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="configMapDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="configMapDetail.data" title="数据" name="data">
          <vue-json-pretty :data="configMapDetail.data" :color="'lightcoral'" />
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="configMapFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getConfigMapDetail, getConfigMapRaw } from '@/api/kubernetes/config/configMap'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import VueJsonPretty from '@/components/vueJsonPretty/index.vue'
export default {
  name: 'ConfigMapDetail',
  components: {
    VueCodeMirror,
    MetaData,
    VueJsonPretty
  },
  setup() {
    const activeNames = ref(['metadata', 'data'])
    const configMapDetail = ref({})
    const configMapFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const configMap = route.query.configMap

    // 加载configMap详情
    const getData = async() => {
      await getConfigMapDetail({ namespace: namespace, configMap: configMap }).then(response => {
        if (response.code === 0) {
          configMapDetail.value = response.data
        }
      })
    }
    getData()

    // 操作
    const viewConfigMap = async() => {
      const result = await getConfigMapRaw({ configMap: configMap, namespace: namespace })
      if (result.code === 0) {
        configMapFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      configMapDetail,
      formatDate,
      dialogFormVisible,
      namespace,
      // 操作
      viewConfigMap,
      configMapFormat
    }
  }
}
</script>
