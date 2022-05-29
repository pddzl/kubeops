<template>
  <div>
    <el-row>
      <el-col :span="6"><node-type /></el-col>
      <el-col :span="6"><node-status /></el-col>
      <el-col :span="6"><namespace-status /></el-col>
      <el-col :span="6"><pod-status /></el-col>
    </el-row>
  </div>
</template>

<script>
import NodeType from './node/type.vue'
import NodeStatus from './node/status.vue'
import NamespaceStatus from './namespace/index.vue'
import PodStatus from './pod/index.vue'
export default {
  name: 'ClusterStatus',
  components: {
    NodeType,
    NodeStatus,
    NamespaceStatus,
    PodStatus
  },
  setup() {
  }
}
</script>

<style lang="scss" scoped>
.cluster {
  display: flex;
}
</style>
