<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewIngress">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="ingressDetail.metadata" title="元数据" name="metadata">
        <meta-data :metadata="ingressDetail.metadata" />
      </el-collapse-item>
      <el-collapse-item v-if="ingressDetail.spec" title="资源信息" name="spec">
        <div class="spec-layout">
          <div v-if="ingressDetail.spec.ingressClassName" class="resource">
            <p>Ingress Class Name</p>
            <div class="item">
              <span class="content">{{ ingressDetail.spec.ingressClassName }}</span>
            </div>
          </div>
          <div v-if="ingressDetail.status.endPoints" class="resource">
            <p>EndPoints</p>
            <div class="item">
              <span v-for="endPoint in ingressDetail.status.endPoints" :key="endPoint">
                {{ endPoint }}
              </span>
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
                    :to="{ name: 'services_detail', query: { service: scope.row.backend.service.name, namespace: namespace } }"
                  >
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
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="ingressFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getIngressDetail, getIngressRaw } from '@/api/kubernetes/ingress'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
export default {
  name: 'ServicesDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    const activeNames = ref(['metadata', 'spec', 'rules'])
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const ingressDetail = ref({})
    const ingressFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const ingress = route.query.ingress

    // 获取daemonSet详情
    const getIngressDetailData = async() => {
      await getIngressDetail({ namespace: namespace, ingress: ingress }).then(response => {
        if (response.code === 0) {
          ingressDetail.value = response.data
        }
      })
    }
    getIngressDetailData()

    // 操作
    const viewIngress = async() => {
      const result = await getIngressRaw({ ingress: ingress, namespace: namespace })
      if (result.code === 0) {
        ingressFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      ingressDetail,
      ingressFormat,
      dialogFormVisible,
      namespace,
      // 格式化日期
      formatDate,
      // 操作
      viewIngress,
      // 分页
      page,
      pageSize,
      total
    }
  }
}
</script>

<style scoped lang="scss">
.el-collapse {
  --el-collapse-header-font-size: 15px;

  .el-collapse-item {
    background-color: #ffffff;
    padding-left: 20px;
  }

  .el-collapse-item:not(:last-child) {
    margin-bottom: 15px;
  }
}
.rules {
  div {
    margin-right: 20px;
  }
  div:not(:last-child) {
    margin-bottom: 20px;
  }
}
</style>
