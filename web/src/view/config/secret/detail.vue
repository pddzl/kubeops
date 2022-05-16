<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewSecret">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="secretDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="secretDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="secretDetail.data" title="数据" name="data">
          <div v-for="(d,k) in secretDetail.data" :key="k" style="margin-bottom: 10px;">
            <div>
              <span style="margin-right: 3px;">{{ k }}</span>
            </div>
            <div class="code-block">
              {{ dataDecode(d) }}
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="secretFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getSecretDetail, getSecretRaw } from '@/api/kubernetes/config/secret'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
export default {
  name: 'SecretDetail',
  components: {
    VueCodeMirror,
    MetaData,
  },
  setup() {
    const activeNames = ref(['metadata', 'data'])
    const secretDetail = ref({})
    const secretFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const secret = route.query.secret

    // 加载secret详情
    const getData = async() => {
      await getSecretDetail({ namespace: namespace, secret: secret }).then(response => {
        if (response.code === 0) {
          secretDetail.value = response.data
        }
      })
    }
    getData()

    const dataDecode = (str) => {
      // const res = Buffer.from(str, 'base64')
      // return res.toString()
      // console.log(str)
      return atob(str)
      // console.log(decodeStr)
    }

    // 操作
    const viewSecret = async() => {
      const result = await getSecretRaw({ secret: secret, namespace: namespace })
      if (result.code === 0) {
        secretFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      secretDetail,
      formatDate,
      dialogFormVisible,
      namespace,
      // 操作
      viewSecret,
      secretFormat,
      dataDecode
    }
  }
}
</script>

<style lang="scss" scoped>
.code-block {
  border-radius: 2px;
  display: block;
  padding: 16px;
  margin-right: 20px;
  white-space: pre-wrap;
  word-break: break-all;
  background-color: #0000001f;
  font-family: Roboto Mono Regular,monospace;
}
.el-icon svg {
  height: 0.8em;
}
</style>
