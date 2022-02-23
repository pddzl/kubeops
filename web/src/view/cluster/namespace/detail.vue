<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="namespaceDetail.metadata" title="元数据" name="1">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>名称</p>
              <span class="content">{{ namespaceDetail.metadata.name }}</span>
            </div>
            <div>
              <p>UID</p>
              <span class="content">{{ namespaceDetail.metadata.uid }}</span>
            </div>
            <div>
              <p>创建时间</p>
              <span class="content">{{ formatDate(namespaceDetail.metadata.creationTimestamp) }}</span>
            </div>
          </div>
        </div>
        <div class="common_show">
          <p>标签:</p>
          <div v-for="(label, index) in namespaceDetail.metadata.labels" :key="index">
            <span class="span-shadow">
              {{ index }}
              <span v-if="label">:</span>
              {{ label }}
            </span>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item title="资源信息" name="2">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>状态</p>
              <span class="content">{{ namespaceDetail.status }}</span>
            </div>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="namespaceDetail.resourceQuotaList" title="资源配额" name="3">
        <p>123</p>
      </el-collapse-item>
      <el-collapse-item v-if="namespaceDetail.resourceLimits && namespaceDetail.resourceLimits.length > 0" title="资源限制" name="4">
        <p>456</p>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getNamespaceDetail } from '@/api/kubernetes/namespace'
import { formatDate } from '@/utils/format'
export default {
  name: 'NamespaceDetail',
  setup() {
    const activeNames = ref(['1', '2', '3', '4'])
    const namespaceDetail = ref({})

    const route = useRoute()
    const namespace = route.query.name

    // 加载namespace详情
    const getData = async () => {
      await getNamespaceDetail({ name: namespace }).then(response => {
        if (response.code === 0) {
          namespaceDetail.value = response.data
        }
      })
    }
    getData()

    return {
      activeNames,
      namespaceDetail,
      formatDate
    }
  }
}
</script>

<style lang="scss" scoped>
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
.address {
  span:not(:last-child) {
    margin-right: 5px;
  }
}
</style>
