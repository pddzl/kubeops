<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewConfigMap">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
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
import { getConfigMapDetail, getConfigMapRaw, deleteConfigMap } from '@/api/kubernetes/config/configMap'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import VueJsonPretty from '@/components/vueJsonPretty/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'ConfigMapDetail',
  components: {
    VueCodeMirror,
    MetaData,
    VueJsonPretty
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'data'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const configMap = route.query.configMap

    // 获取configMap详情
    const configMapDetail = ref({})

    const getData = async() => {
      await getConfigMapDetail({ namespace: namespace, configMap: configMap }).then(response => {
        if (response.code === 0) {
          configMapDetail.value = response.data
        }
      })
    }
    getData()

    // 编辑
    const configMapFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewConfigMap = async() => {
      const result = await getConfigMapRaw({ configMap: configMap, namespace: namespace })
      if (result.code === 0) {
        configMapFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该ConfigMap, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteConfigMap({ namespace: namespace, configMap: configMap })
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
      configMapDetail,
      // time format
      formatDate,
      // 编辑
      configMapFormat,
      dialogFormVisible,
      viewConfigMap,
      // 删除
      deleteFunc
    }
  }
}
</script>
