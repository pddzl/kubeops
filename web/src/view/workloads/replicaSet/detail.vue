<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain>查看</el-button>
          <el-button icon="expand" size="small" type="warning" plain>伸缩</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="replicaSetDetail.objectMeta" title="元数据" name="1">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>名称</p>
              <span class="content">{{ replicaSetDetail.objectMeta.name }}</span>
            </div>
            <div>
              <p>命名空间</p>
              <span class="content">{{ replicaSetDetail.objectMeta.namespace }}</span>
            </div>
            <div>
              <p>UID</p>
              <span class="content">{{ replicaSetDetail.objectMeta.uid }}</span>
            </div>
            <div>
              <p>创建时间</p>
              <span class="content">{{ formatDate(replicaSetDetail.objectMeta.creationTimestamp) }}</span>
            </div>
          </div>
        </div>
        <div class="common_show">
          <p>标签:</p>
          <div v-for="(label, index) in replicaSetDetail.objectMeta.labels" :key="index">
            <span class="span-shadow">
              {{ index }}
              <span v-if="label">:</span>
              {{ label }}
            </span>
          </div>
        </div>
        <div class="common_show">
          <p>注释:</p>
          <div v-for="(label, index) in replicaSetDetail.objectMeta.annotations" :key="index">
            <span class="span-shadow">
              {{ index }}
              <span v-if="label">:</span>
              {{ label }}
            </span>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item title="资源信息" name="2">
        <div v-if="replicaSetDetail.selector" class="common_show">
          <p>选择器</p>
          <div v-for="(label, index) in replicaSetDetail.selector.matchLabels" :key="index">
            <span class="span-shadow">
              {{ index }}
              <span v-if="label">:</span>
              {{ label }}
            </span>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="replicaSetDetail.podInfo" title="Pods状态" name="3">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>current</p>
              <span class="content">{{ replicaSetDetail.podInfo.current }}</span>
            </div>
            <div>
              <p>desired</p>
              <span class="content">{{ replicaSetDetail.podInfo.desired }}</span>
            </div>
            <div>
              <p>running</p>
              <span class="content">{{ replicaSetDetail.podInfo.running }}</span>
            </div>
            <div>
              <p>pending</p>
              <span class="content">{{ replicaSetDetail.podInfo.pending }}</span>
            </div>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item title="Pods" name="4">
        <div style="padding-right: 20px;">
          <el-table :data="replicaSetPods">
            <el-table-column label="名称" prop="objectMeta.name" min-width="120">
              <template #default="scope">
                <router-link
                  :to="{ name: 'pod_detail', query: { pod: scope.row.objectMeta.name, namespace: scope.row.objectMeta.namespace } }"
                >
                  <el-link type="primary" :underline="false">{{ scope.row.objectMeta.name }}</el-link>
                </router-link>
              </template>
            </el-table-column>
            <el-table-column label="命名空间" prop="objectMeta.namespace" />
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
              <template #default="scope">{{ formatDate(scope.row.objectMeta.creationTimestamp) }}</template>
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
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getReplicaSetDetail, getReplicaSetPods, getReplicaSetServices } from '@/api/kubernetes/replicaSet'
import { statusPodFilter } from '@/mixin/filter.js'
import { formatDate } from '@/utils/format'
export default {
  name: 'NodeDetail',
  setup() {
    const activeNames = ref(['1', '2', '3', '4', '5'])
    const replicaSetDetail = ref({})
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const replicaSetPods = ref([])
    const replicaSetServices = ref([])

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
      // filter
      statusPodFilter,
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
.interval {
  span:not(:last-child) {
    margin-right: 5px;
  }
}
</style>
