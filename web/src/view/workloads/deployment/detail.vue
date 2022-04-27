<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewDeployment">查看</el-button>
          <el-button icon="expand" size="small" type="warning" plain>伸缩</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="deploymentDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="deploymentDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="deploymentDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ deploymentDetail.spec.replicas }}</span>
              </div>
              <div class="item">
                <p>Selector</p>
                <div class="content">
                  <span v-for="(label, index) in deploymentDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>Strategy</p>
                <span class="content">{{ deploymentDetail.spec.strategy.type }}</span>
              </div>
              <div v-if="deploymentDetail.spec.strategy.type === 'RollingUpdate'" class="item">
                <p>maxUnavailable</p>
                <span class="content">{{ deploymentDetail.spec.strategy.rollingUpdate.maxUnavailable }}</span>
              </div>
              <div v-if="deploymentDetail.spec.strategy.type === 'RollingUpdate'" class="item">
                <p>maxSurge</p>
                <span class="content">{{ deploymentDetail.spec.strategy.rollingUpdate.maxUnavailable }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="deploymentDetail.status" title="状态" name="status">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ deploymentDetail.status.replicas }}</span>
              </div>
              <div class="item">
                <p>updatedReplicas</p>
                <span class="content">{{ deploymentDetail.status.updatedReplicas }}</span>
              </div>
              <div class="item">
                <p>readyReplicas</p>
                <span class="content">{{ deploymentDetail.status.readyReplicas }}</span>
              </div>
              <div class="item">
                <p>availableReplicas</p>
                <span class="content">{{ deploymentDetail.status.availableReplicas }}</span>
              </div>
            </div>
            <div style="margin: 20px 20px 0px 0px">
              <el-table :data="deploymentDetail.status.conditions">
                <el-table-column label="类别" prop="type" />
                <el-table-column label="状态" prop="status">
                  <template #default="scope">
                    <el-tag :type="statusRsFilter(scope.row.status)" size="small">
                      {{ scope.row.status }}
                    </el-tag>
                  </template>
                </el-table-column>
                <el-table-column label="上次检测时间">
                  <template #default="scope">
                    {{ formatDate(scope.row.lastUpdateTime) }}
                  </template>
                </el-table-column>
                <el-table-column label="上次迁移时间">
                  <template #default="scope">
                    {{ formatDate(scope.row.lastTransitionTime) }}
                  </template>
                </el-table-column>
                <el-table-column label="原因" prop="reason" />
                <el-table-column label="信息" prop="message" :show-overflow-tooltip="true" />
              </el-table>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item title="ReplicaSet" name="replicaSet">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>名称</p>
                <div class="content">
                  <router-link
                    :to="{ name: 'replicaSet_detail', query: { replicaSet: newReplicaSet.name, namespace: newReplicaSet.namespace } }"
                  >
                    <el-link type="primary" :underline="false">{{ newReplicaSet.name }}
                    </el-link>
                  </router-link>
                </div>
              </div>
              <div class="item">
                <p>命名空间</p>
                <span class="content">{{ newReplicaSet.namespace }}</span>
              </div>
              <div class="item">
                <p>Replicas</p>
                <span class="content">{{ newReplicaSet.replicas }}</span>
              </div>
              <div class="item">
                <p>创建时间</p>
                <span class="content">{{ formatDate(newReplicaSet.creationTimestamp) }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>标签:</p>
                <div class="content">
                  <span v-for="(label, index) in newReplicaSet.labels" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="deploymentFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getDeploymentDetail, getDeploymentRaw, getNewReplicaSet } from '@/api/kubernetes/deployment'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import { formatDate } from '@/utils/format'
import { statusRsFilter } from '@/mixin/filter.js'
export default {
  name: 'DeploymentDetail',
  components: {
    VueCodeMirror,
    MetaData
  },
  setup() {
    const activeNames = ref(['metadata', 'spec', 'status', 'replicaSet'])
    const deploymentDetail = ref({})
    const newReplicaSet = ref({})
    const deploymentFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const deployment = route.query.deployment

    // 获取deployment详情
    const getDeploymentDetailData = async() => {
      await getDeploymentDetail({ namespace: namespace, deployment: deployment }).then(response => {
        if (response.code === 0) {
          deploymentDetail.value = response.data
        }
      })
    }
    getDeploymentDetailData()

    // 获取deployment关联的replicaset
    const getNewReplicaSetData = async() => {
      await getNewReplicaSet({ namespace: namespace, deployment: deployment }).then(response => {
        if (response.code === 0) {
          newReplicaSet.value = response.data
        }
      })
    }
    getNewReplicaSetData()

    // 操作
    const viewDeployment = async() => {
      const result = await getDeploymentRaw({ deployment: deployment, namespace: namespace })
      if (result.code === 0) {
        deploymentFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      deploymentDetail,
      deploymentFormat,
      dialogFormVisible,
      newReplicaSet,
      // status filter
      statusRsFilter,
      // 格式化日期
      formatDate,
      // 操作
      viewDeployment
    }
  }
}
</script>
