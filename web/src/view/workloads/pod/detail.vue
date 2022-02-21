<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="podDetail.metadata" title="元数据" name="1">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>名称</p>
              <span class="content">{{ podDetail.metadata.name }}</span>
            </div>
            <div>
              <p>命名空间</p>
              <span class="content">{{ podDetail.metadata.namespace }}</span>
            </div>
            <div>
              <p>UID</p>
              <span class="content">{{ podDetail.metadata.uid }}</span>
            </div>
            <div>
              <p>创建时间</p>
              <span class="content">{{ formatDate(podDetail.metadata.createTimestamp) }}</span>
            </div>
          </div>
        </div>
        <div class="common_show">
          <p>标签</p>
          <span
            v-for="(label, index) in podDetail.metadata.labels"
            :key="index"
            class="span-shadow"
          >{{ index }}: {{ label }}</span>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="podDetail.resource_info" title="资源信息" name="2">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>节点</p>
              <span class="content">{{ podDetail.resource_info.node }}</span>
            </div>
            <div>
              <p>状态</p>
              <span class="content">{{ podDetail.resource_info.phase }}</span>
            </div>
            <div>
              <p>IP</p>
              <span class="content">{{ podDetail.resource_info.ip }}</span>
            </div>
            <div>
              <p>服务质量</p>
              <span class="content">{{ podDetail.resource_info.qosClass }}</span>
            </div>
            <div>
              <p>重启策略</p>
              <span class="content">{{ podDetail.resource_info.restartPolicy }}</span>
            </div>
            <div>
              <p>重启次数</p>
              <span class="content">{{ podDetail.resource_info.restarts }}</span>
            </div>
            <div>
              <p>服务账号</p>
              <span class="content">{{ podDetail.resource_info.serviceAccount }}</span>
            </div>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="podDetail.conditions" title="状态" name="3">
        <div class="table">
          <el-table :data="podDetail.conditions">
            <el-table-column prop="type" label="类别" min-width="120" />
            <el-table-column prop="status" label="状态" />
            <el-table-column label="最后检查时间" min-width="200">
              <template #default="scope">
                <span v-if="scope.row.lastProbeTime">{{ formatDate(scope.row.lastProbeTime) }}</span>
                <span v-else>-</span>
              </template>
            </el-table-column>
            <el-table-column label="最后迁移时间" min-width="200">
              <template #default="scope">{{ formatDate(scope.row.lastTransitionTime) }}</template>
            </el-table-column>
          </el-table>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="podDetail.ownerReferences && Object.keys(podDetail.ownerReferences).length > 0" title="控制器" name="4">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>名称</p>
              <span class="content">{{ podDetail.ownerReferences.name }}</span>
            </div>
            <div>
              <p>类型</p>
              <span class="content">{{ podDetail.ownerReferences.kind }}</span>
            </div>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="podDetail.containers && podDetail.containers.length > 0" title="容器" name="5">
        <container :data="podDetail.containers" />
      </el-collapse-item>
      <el-collapse-item v-if="podDetail.initContainers && podDetail.initContainers.length > 0" title="初始容器" name="6">
        <container :data="podDetail.initContainers" />
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getPodDetail } from '@/api/kubernetes/pod'
import { formatDate } from '@/utils/format'
import Container from './components/container.vue'
export default {
  name: 'PodDetail',
  components: {
    Container
  },
  setup() {
    const activeNames = ref(['1', '2', '3', '4', '5', '6'])
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
      formatDate,
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
  margin-right: 10px;
}
.table {
  margin-right: 20px;
}
</style>
