<template>
  <div>
    <div class="detail-operation">
      <el-button icon="view" size="small" type="primary" plain @click="viewSecret">查看</el-button>
      <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
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
import { getSecretDetail, getSecretRaw, deleteSecret } from '@/api/kubernetes/config/secret'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'SecretDetail',
  components: {
    VueCodeMirror,
    MetaData,
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'data'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const secret = route.query.secret

    // secret detail
    const secretDetail = ref({})

    const getData = async() => {
      await getSecretDetail({ namespace: namespace, secret: secret }).then(response => {
        if (response.code === 0) {
          secretDetail.value = response.data
        }
      })
    }
    getData()

    // 编辑
    const secretFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewSecret = async() => {
      const result = await getSecretRaw({ secret: secret, namespace: namespace })
      if (result.code === 0) {
        secretFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该Secret, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteSecret({ namespace: namespace, secret: secret })
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '删除成功!'
            })
          }
        })
    }

    // 解码
    const dataDecode = (str) => {
      return atob(str)
    }

    return {
      // 折叠面板
      activeNames,
      // data
      namespace,
      secretDetail,
      // time format
      formatDate,
      // 编辑
      dialogFormVisible,
      secretFormat,
      viewSecret,
      // 删除
      deleteFunc,
      // 解码
      dataDecode,
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
