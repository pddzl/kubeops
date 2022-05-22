<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="viewReplicaSet">查看</el-button>
          <el-button icon="expand" size="small" type="warning" plain :disabled="namespace === 'kube-system'" @click="openScaleDialog">伸缩
          </el-button>
          <el-button icon="delete" size="small" type="danger" plain :disabled="namespace === 'kube-system'">删除
          </el-button>
        </el-affix>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="replicaSetDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="replicaSetDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="replicaSetDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ replicaSetDetail.spec.replicas }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>Selector:</p>
                <div class="content">
                  <span v-for="(label, index) in replicaSetDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="replicaSetDetail.status" title="状态" name="status">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ replicaSetDetail.status.replicas }}</span>
              </div>
              <div class="item">
                <p>fullyLabeledReplicas</p>
                <span class="content">{{ replicaSetDetail.status.fullyLabeledReplicas }}</span>
              </div>
              <div class="item">
                <p>readyReplicas</p>
                <span class="content">{{ replicaSetDetail.status.readyReplicas }}</span>
              </div>
              <div class="item">
                <p>availableReplicas</p>
                <span class="content">{{ replicaSetDetail.status.availableReplicas }}</span>
              </div>
            </div>
          </div>
          <div v-if="replicaSetDetail.status.conditions" style="margin-top: 20px; margin-right: 20px;">
            <el-table :data="replicaSetDetail.status.conditions">
              <el-table-column label="类别" prop="type" />
              <el-table-column label="状态" prop="status">
                <template #default="scope">
                  <el-tag :type="statusRsFilter(scope.row.status)" size="small">
                    {{ scope.row.status }}
                  </el-tag>
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
        </el-collapse-item>
        <el-collapse-item title="Pods" name="pods">
          <div style="margin-right: 20px;">
            <el-table :data="replicaSetPods">
              <el-table-column label="名称" prop="metadata.name" min-width="120">
                <template #default="scope">
                  <router-link
                    :to="{ name: 'pod_detail', query: { pod: scope.row.metadata.name, namespace: scope.row.metadata.namespace } }"
                  >
                    <el-link type="primary" :underline="false">{{ scope.row.metadata.name }}</el-link>
                  </router-link>
                </template>
              </el-table-column>
              <el-table-column label="命名空间" prop="metadata.namespace" />
              <el-table-column label="节点" prop="nodeName" />
              <el-table-column label="状态">
                <template #default="scope">
                  <el-tag :type="statusPodFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="创建时间">
                <template #default="scope">{{ formatDate(scope.row.metadata.creationTimestamp) }}</template>
              </el-table-column>
            </el-table>
            <div class="gva-pagination">
              <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :total="total"
                layout="total, sizes, prev, pager, next, jumper"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
              />
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item title="Services" name="services">
          <div style="margin-right: 20px;">
            <el-table :data="replicaSetServices">
              <el-table-column label="名称" prop="metadata.name" min-width="120">
                <template #default="scope">
                  <router-link
                    :to="{ name: 'services_detail', query: { service: scope.row.metadata.name, namespace: namespace } }"
                  >
                    <el-link type="primary" :underline="false">{{ scope.row.metadata.name }}</el-link>
                  </router-link>
                </template>
              </el-table-column>
              <el-table-column label="命名空间" prop="metadata.namespace" />
              <el-table-column label="类型" prop="spec.type" />
              <el-table-column label="集群IP" prop="spec.clusterIP" />
              <el-table-column label="创建时间">
                <template #default="scope">{{ formatDate(scope.row.metadata.creationTimestamp) }}</template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <!-- 查看编排对话框 -->
    <el-dialog v-model="dialogViewVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="replicaSetFormat" :readOnly="true" />
    </el-dialog>
    <!-- 伸缩对话框 -->
    <el-dialog v-model="dialogScaleVisible" title="伸缩" width="55%" center>
      <p style="font-weight: bold;">ReplicaSet {{ replicaSet }} will be updated to reflect the
        desired replicas count.</p>
      <div style="margin: 25px 0 25px 0px;">
        <span style="margin-right: 10px;">Desired Replicas:</span>
        <el-input-number v-model="desiredNum" :min="0" :max="50" style="margin-right: 20px;" />
        <span style="margin-right: 10px;">Actual Replicas: </span>
        <el-input-number v-model="ActualNum" disabled />
      </div>
      <warning-bar :title="warningTitle" />
      <template #footer>
        <el-button @click="closeScaleDialog">取消</el-button>
        <el-button type="primary" @click="scaleFunc">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { getReplicaSetDetail, getReplicaSetPods, getReplicaSetServices, getReplicaSetRaw } from '@/api/kubernetes/replicaSet'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { statusPodFilter, statusRsFilter } from '@/mixin/filter.js'
import { scale } from '@/api/kubernetes/scale'
import { formatDate } from '@/utils/format'
import MetaData from '@/components/kubernetes/detail/metadata.vue'
import warningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage } from 'element-plus'
export default {
  name: 'ReplicaSetDetail',
  components: {
    VueCodeMirror,
    MetaData,
    warningBar
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'spec', 'status', 'pods', 'services'])

    // 路由
    const route = useRoute()
    const namespace = route.query.namespace
    const replicaSet = route.query.replicaSet

    // 加载replicaSet详情
    const replicaSetDetail = ref({})

    const getData = async() => {
      await getReplicaSetDetail({ namespace: namespace, replicaSet: replicaSet }).then(response => {
        if (response.code === 0) {
          replicaSetDetail.value = response.data
        }
      })
    }
    getData()

    // 加载replicaSet关联pods
    const replicaSetPods = ref([])
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)

    const getReplicaSetPodsData = async() => {
      const table = await getReplicaSetPods({ page: page.value, pageSize: pageSize.value, namespace: namespace, replicaSet: replicaSet })
      if (table.code === 0) {
        replicaSetPods.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getReplicaSetPodsData()

    // 加载replicaSet关联services
    const replicaSetServices = ref([])

    const getReplicaSetServicesData = async() => {
      const table = await getReplicaSetServices({ namespace: namespace, replicaSet: replicaSet })
      if (table.code === 0) {
        replicaSetServices.value = table.data
      }
    }
    getReplicaSetServicesData()

    // 查看编排
    const dialogViewVisible = ref(false)
    const replicaSetFormat = ref({})

    const viewReplicaSet = async() => {
      const result = await getReplicaSetRaw({ replicaSet: replicaSet, namespace: namespace })
      if (result.code === 0) {
        replicaSetFormat.value = JSON.stringify(result.data)
      }
      dialogViewVisible.value = true
    }

    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getReplicaSetPodsData()
    }

    const handleCurrentChange = (val) => {
      page.value = val
      getReplicaSetPodsData()
    }

    // 伸缩
    const dialogScaleVisible = ref(false)
    const warningTitle = ref('')
    const desiredNum = ref(0)
    const ActualNum = ref(0)

    // -> 打开对话框
    const openScaleDialog = () => {
      desiredNum.value = replicaSetDetail.value.status.replicas
      ActualNum.value = replicaSetDetail.value.status.availableReplicas
      warningTitle.value = `This action is equivalent to: kubectl scale -n ${namespace} replicaset ${replicaSet} --replicas=${ActualNum.value}`
      dialogScaleVisible.value = true
    }

    watch(desiredNum, (val) => {
      warningTitle.value = `This action is equivalent to: kubectl scale -n ${namespace} replicaset ${replicaSet} --replicas=${val}`
    })

    // -> 关闭对话框
    const closeScaleDialog = () => {
      dialogScaleVisible.value = false
    }

    // -> 操作
    const scaleFunc = async() => {
      const res = await scale({ namespace: namespace, name: replicaSet, kind: 'replicaSet', num: desiredNum.value })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '伸缩成功',
          showClose: true
        })
        // 查询刷新数据
        getData()
        getReplicaSetPodsData()
        getReplicaSetServicesData()
      }
      dialogScaleVisible.value = false
    }

    return {
      // 折叠面板
      activeNames,
      // replicaSet相关
      namespace,
      replicaSet,
      replicaSetDetail,
      replicaSetPods,
      replicaSetServices,
      // time format
      formatDate,
      // filter
      statusPodFilter,
      statusRsFilter,
      // 查看编排
      dialogViewVisible,
      viewReplicaSet,
      replicaSetFormat,
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange,
      // 伸缩
      dialogScaleVisible,
      warningTitle,
      desiredNum,
      ActualNum,
      openScaleDialog,
      closeScaleDialog,
      scaleFunc
    }
  }
}
</script>
