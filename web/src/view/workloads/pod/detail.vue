<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="podDetail.metadata" class="object_meta" title="元数据" name="1">
        <span style="font-size: 14px; margin-right: 150px;">名称: {{ podDetail.metadata.name }}</span>
        <span style="font-size: 14px; margin-right: 150px;">命名空间: {{ podDetail.metadata.namespace }}</span>
        <span style="font-size: 14px; margin-right: 150px;">UID: {{ podDetail.metadata.uid }}</span>
        <span style="font-size: 14px">创建时间: {{ formatDate(podDetail.metadata.createTimestamp) }}</span>
        <p>标签:</p>
        <span v-for="(label, index) in podDetail.metadata.labels" :key="index" class="span-shadow">
          {{ index }}: {{ label }}
        </span>
        <div v-if="podDetail.metadata.ownerReferences.controller">
          <p>控制器:</p>
          <span
            style="font-size: 14px; margin-right: 150px;"
          >名称: {{ podDetail.metadata.ownerReferences.name }}</span>
          <span
            style="font-size: 14px; margin-right: 150px;"
          >类别: {{ podDetail.metadata.ownerReferences.kind }}</span>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="podDetail.resource_info" title="资源信息" name="2">
        <el-descriptions direction="vertical" :column="7">
          <el-descriptions-item label="节点">{{ podDetail.resource_info.node }}</el-descriptions-item>
          <el-descriptions-item label="状态">{{ podDetail.resource_info.phase }}</el-descriptions-item>
          <el-descriptions-item label="IP">{{ podDetail.resource_info.ip }}</el-descriptions-item>
          <el-descriptions-item label="服务质量">{{ podDetail.resource_info.qosClass }}</el-descriptions-item>
          <el-descriptions-item label="重启策略">{{ podDetail.resource_info.restartPolicy }}</el-descriptions-item>
          <el-descriptions-item label="重启次数">{{ podDetail.resource_info.restarts }}</el-descriptions-item>
          <el-descriptions-item label="服务账号">{{ podDetail.resource_info.serviceAccount }}</el-descriptions-item>
        </el-descriptions>
      </el-collapse-item>
      <el-collapse-item v-if="podDetail.conditions" title="状态" name="3">
        <el-table :data="podDetail.conditions">
          <el-table-column prop="type" label="类别" min-width="120" />
          <el-table-column prop="status" label="状态" />
          <el-table-column label="最后检查时间" min-width="200">
            <template #default="scope">
              <span v-if="scope.row.lastProbeTime">
                {{ formatDate(scope.row.lastProbeTime) }}
              </span>
              <span v-else>
                -
              </span>
            </template>
          </el-table-column>
          <el-table-column label="最后迁移时间" min-width="200">
            <template #default="scope">{{ formatDate(scope.row.lastTransitionTime) }}</template>
          </el-table-column>
        </el-table>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getPodDetail } from '@/api/kubernetes/pod'
import { formatDate } from '@/utils/format'
export default {
  name: 'PodDetail',
  setup() {
    const activeNames = ref(['1', '2', '3'])
    const podDetail = ref({})

    const route = useRoute()
    const pod = route.query.pod
    const namespace = route.query.namespace

    // 加载pod详情
    const getData = async() => {
      await getPodDetail({ pod: pod, namespace: namespace }).then(response => {
        if (response.code === 0) {
          podDetail.value = response.data
        }
      })
    }
    getData()

    return {
      activeNames,
      podDetail,
      formatDate
    }
  }
}
</script>

<style scoped lang="scss">
.object_meta {
  p {
    font-size: 14px;
    margin-top: 15px;
  }
}
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
  margin-right: 10px;
}
</style>
