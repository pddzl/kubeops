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
              <span class="content">
                <el-tag
                  :type="statusNsFilter(namespaceDetail.status)"
                  size="small"
                >{{ namespaceDetail.status }}</el-tag>
              </span>
            </div>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item
        v-if="namespaceDetail.resourceQuotaList && namespaceDetail.resourceQuotaList.length > 0"
        title="资源配额"
        name="3"
      >
        <div class="table">
          <el-table :data="namespaceDetail.resourceQuotaList">
            <el-table-column label="名称" align="center">
              <template #default="scope">
                <el-tag size="small" effect="plain">{{ scope.row.objectMeta.name }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="limits" align="center">
              <el-table-column label="cpu used" align="center" prop="statusList.limits_cpu.used" />
              <el-table-column label="cpu hard" align="center" prop="statusList.limits_cpu.hard" />
              <el-table-column
                label="memory used"
                align="center"
                prop="statusList.limits_memory.used"
              />
              <el-table-column
                label="memory hard"
                align="center"
                prop="statusList.limits_memory.hard"
              />
            </el-table-column>
            <el-table-column label="requests" align="center">
              <el-table-column label="cpu used" align="center" prop="statusList.requests_cpu.used" />
              <el-table-column label="cpu hard" align="center" prop="statusList.requests_cpu.hard" />
              <el-table-column
                label="memory used"
                align="center"
                prop="statusList.requests_memory.used"
              />
              <el-table-column
                label="memory hard"
                align="center"
                prop="statusList.requests_memory.hard"
              />
            </el-table-column>
          </el-table>
        </div>
      </el-collapse-item>
      <el-collapse-item
        v-if="namespaceDetail.resourceLimits && namespaceDetail.resourceLimits.length > 0"
        title="资源限制"
        name="4"
      >
        <div v-for="(rLimit, keyRl) in namespaceDetail.resourceLimits" :key="keyRl" class="table">
          <div class="name">
            <el-tag size="small">{{ rLimit.name }}</el-tag>
          </div>
          <el-table :data="rLimit.limits">
            <el-table-column label="类型" align="center" prop="type" />
            <el-table-column label="max" align="center">
              <el-table-column label="cpu" align="center" prop="max.cpu" />
              <el-table-column label="memory" align="center" prop="max.memory" />
            </el-table-column>
            <el-table-column label="min" align="center">
              <el-table-column label="cpu" align="center" prop="min.cpu" />
              <el-table-column label="memory" align="center" prop="min.memory" />
            </el-table-column>
            <el-table-column label="default" align="center">
              <el-table-column label="cpu" align="center" prop="default.cpu" />
              <el-table-column label="memory" align="center" prop="default.memory" />
            </el-table-column>
            <el-table-column label="defaultRequest" align="center">
              <el-table-column label="cpu" align="center" prop="defaultRequest.cpu" />
              <el-table-column label="memory" align="center" prop="defaultRequest.memory" />
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
import { formatDate } from '@/utils/format'
import { statusNsFilter } from '@/mixin/filter'
import { getNamespaceDetail } from '@/api/kubernetes/namespace'
export default {
  name: 'NamespaceDetail',
  setup() {
    const activeNames = ref(['1', '2', '3', '4'])
    const namespaceDetail = ref({})

    const route = useRoute()
    const namespace = route.query.name

    // 加载namespace详情
    const getData = async() => {
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
      formatDate,
      statusNsFilter
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
.table {
  margin-right: 20px;
}
.name {
  margin-bottom: 10px;
}
</style>
