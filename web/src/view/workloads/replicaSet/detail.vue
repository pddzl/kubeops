<template>
  <div>
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
    </el-collapse>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getReplicaSetDetail } from '@/api/kubernetes/replicaSet'
import { statusPodFilter } from '@/mixin/filter.js'
import { formatDate } from '@/utils/format'
export default {
  name: 'NodeDetail',
  setup() {
    const activeNames = ref(['1', '2', '3'])
    const replicaSetDetail = ref({})

    const route = useRoute()
    const namespace = route.query.namespace
    const replicaSet = route.query.replicaSet

    // 加载node详情
    const getData = async() => {
      await getReplicaSetDetail({ namespace: namespace, replicaSet: replicaSet }).then(response => {
        if (response.code === 0) {
          replicaSetDetail.value = response.data
        }
      })
    }
    getData()

    return {
      activeNames,
      replicaSetDetail,
      formatDate,
      // filter
      statusPodFilter
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
