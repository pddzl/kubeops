<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="editReplicaSet">查看</el-button>
          <el-button icon="expand" size="small" type="warning" plain>伸缩</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="replicaSetDetail.metadata" title="元数据" name="1">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>名称</p>
              <span class="content">{{ replicaSetDetail.metadata.name }}</span>
            </div>
            <div>
              <p>命名空间</p>
              <span class="content">{{ replicaSetDetail.metadata.namespace }}</span>
            </div>
            <div>
              <p>UID</p>
              <span class="content">{{ replicaSetDetail.metadata.uid }}</span>
            </div>
            <div>
              <p>创建时间</p>
              <span class="content">{{ formatDate(replicaSetDetail.metadata.creationTimestamp) }}</span>
            </div>
          </div>
        </div>
        <div class="common_show">
          <p>标签:</p>
          <span v-for="(label, index) in replicaSetDetail.metadata.labels" :key="index" class="span-shadow">
            {{ index }}
            <span v-if="label">:</span>
            {{ label }}
          </span>
        </div>
        <div class="common_show">
          <p>注释:</p>
          <div style="color: lightcoral;">
            <!-- eslint-disable-next-line vue/attribute-hyphenation -->
            <vue-json-pretty :data="replicaSetDetail.metadata.annotations" :deep="2" :showLine="true" />
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="replicaSetDetail.spec" title="资源信息" name="2">
        <div class="row_mine" style="margin-bottom: 10px;">
          <div class="row_context">
            <div>
              <p>replicas</p>
              <span class="content">{{ replicaSetDetail.spec.replicas }}</span>
            </div>
          </div>
        </div>
        <div class="common_show">
          <p>Selector:</p>
          <span v-for="(label, index) in replicaSetDetail.spec.selector" :key="index" class="span-shadow">
            {{ index }}
            <span v-if="label">:</span>
            {{ label }}
          </span>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="replicaSetDetail.status" title="状态" name="3">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>replicas</p>
              <span class="content">{{ replicaSetDetail.status.replicas }}</span>
            </div>
            <div>
              <p>fullyLabeledReplicas</p>
              <span class="content">{{ replicaSetDetail.status.fullyLabeledReplicas }}</span>
            </div>
            <div>
              <p>readyReplicas</p>
              <span class="content">{{ replicaSetDetail.status.readyReplicas }}</span>
            </div>
            <div>
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
      <el-collapse-item title="Pods" name="4">
        <div style="padding-right: 20px;">
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
                <el-tag
                  :type="statusPodFilter(scope.row.status)"
                  size="small"
                >{{ scope.row.status }}</el-tag>
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
      <el-collapse-item title="Services" name="5">
        <div style="padding-right: 20px;">
          <el-table :data="replicaSetServices">
            <el-table-column label="名称" prop="metadata.name" min-width="120" />
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
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="replicaSetFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getReplicaSetDetail, getReplicaSetPods, getReplicaSetServices, getReplicaSetRaw } from '@/api/kubernetes/replicaSet'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { statusPodFilter } from '@/mixin/filter.js'
import { formatDate } from '@/utils/format'
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'
export default {
  name: 'ReplicaSetDetail',
  components: {
    VueCodeMirror,
    VueJsonPretty
  },
  setup() {
    const activeNames = ref(['1', '2', '3', '4', '5'])
    const replicaSetDetail = ref({})
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const replicaSetPods = ref([])
    const replicaSetServices = ref([])
    const replicaSetFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const replicaSet = route.query.replicaSet

    // 加载replicaSet详情
    const getData = async() => {
      await getReplicaSetDetail({ namespace: namespace, replicaSet: replicaSet }).then(response => {
        if (response.code === 0) {
          replicaSetDetail.value = response.data
        }
      })
    }
    getData()

    // 加载关联pods
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

    // 加载关联services
    const getReplicaSetServicesData = async() => {
      const table = await getReplicaSetServices({ namespace: namespace, replicaSet: replicaSet })
      if (table.code === 0) {
        replicaSetServices.value = table.data
      }
    }
    getReplicaSetServicesData()

    // 操作
    const editReplicaSet = async() => {
      const result = await getReplicaSetRaw({ replicaSet: replicaSet, namespace: namespace })
      if (result.code === 0) {
        replicaSetFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
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

    return {
      // 响应式数据
      activeNames,
      replicaSetDetail,
      replicaSetPods,
      replicaSetServices,
      formatDate,
      dialogFormVisible,
      // filter
      statusPodFilter,
      // 操作
      editReplicaSet,
      replicaSetFormat,
      // 分页
      page,
      pageSize,
      total,
      handleSizeChange,
      handleCurrentChange
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
.span-shadow:not(:last-child) {
  margin-right: 5px;
}
</style>
