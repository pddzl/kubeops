<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewIngress">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="ingressDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="ingressDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="ingressDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div v-if="ingressDetail.spec.ingressClassName" class="row">
              <div class="item">
                <p>Ingress Class Name</p>
                <span class="content">{{ ingressDetail.spec.ingressClassName }}</span>
              </div>
            </div>
            <div v-if="ingressDetail.status.endPoints" class="row">
              <div class="item">
                <p>EndPoints</p>
                <div class="content">
                  <span v-for="endPoint in ingressDetail.status.endPoints" :key="endPoint" class="shadow">
                    {{ endPoint }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="ingressDetail.spec" title="规则" name="rules" class="rules">
          <div v-for="h in ingressDetail.spec.rules" :key="h">
            <el-table :data="h.http.paths">
              <el-table-column label="主机" prop="h.host" />
              <el-table-column label="路径" prop="path" />
              <el-table-column label="路径类型" prop="pathType" />
              <el-table-column label="后端" align="center">
                <el-table-column label="服务名称">
                  <template #default="scope">
                    <router-link
                      :to="{ name: 'services_detail', query: { service: scope.row.backend.service.name, namespace: namespace } }">
                      <el-link type="primary" :underline="false">{{ scope.row.backend.service.name }}</el-link>
                    </router-link>
                  </template>
                </el-table-column>
                <el-table-column label="服务端口" prop="backend.service.port.number" />
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="ingressFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getIngressDetail, getIngressRaw, deleteIngress } from '@/api/kubernetes/ingress'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'ServicesDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'spec', 'rules'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const ingress = route.query.ingress

    // 获取ingress详情
    const ingressDetail = ref({})

    const getIngressDetailData = async() => {
      await getIngressDetail({ namespace: namespace, ingress: ingress }).then(response => {
        if (response.code === 0) {
          ingressDetail.value = response.data
        }
      })
    }
    getIngressDetailData()

    // 编辑
    const ingressFormat = ref({})
    const dialogFormVisible = ref(false)

    const viewIngress = async() => {
      const result = await getIngressRaw({ ingress: ingress, namespace: namespace })
      if (result.code === 0) {
        ingressFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 删除
    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该Ingress, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteIngress({ namespace: namespace, ingress: ingress })
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
      ingressDetail,
      namespace,
      // 格式化日期
      formatDate,
      // 编辑
      ingressFormat,
      dialogFormVisible,
      viewIngress,
      // 删除
      deleteFunc
    }
  }
}
</script>

<style scoped lang="scss">
.rules {
  div {
    margin-right: 20px;
  }
  div:not(:last-child) {
    margin-bottom: 20px;
  }
}
</style>
