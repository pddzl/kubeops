<template>
  <el-table :data="pods">
    <el-table-column label="名称" min-width="140">
      <template #default="scope">
        <router-link
          :to="{ name: 'pod_detail', query: { pod: scope.row.metadata.name, namespace: scope.row.metadata.namespace } }"
        >
          <el-link type="primary" :underline="false">{{ scope.row.metadata.name }}</el-link>
        </router-link>
      </template>
    </el-table-column>
    <el-table-column prop="status" label="状态" min-width="60">
      <template #default="scope">
        <el-tag :type="statusPodFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="node" label="节点" min-width="80" />
    <el-table-column prop="metadata.namespace" label="命名空间" min-width="80" />
    <el-table-column label="创建时间" width="200">
      <template #default="scope">{{ formatDate(scope.row.metadata.creationTimestamp) }}</template>
    </el-table-column>
    <el-table-column fixed="right" label="操作" width="240">
      <template #default>
        <el-button icon="tickets" type="text" size="small">日志</el-button>
        <el-button icon="ArrowRight" type="text" size="small">终端</el-button>
        <el-button icon="delete" type="text" size="small">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import { statusPodFilter } from '@/mixin/filter.js'
import { formatDate } from '@/utils/format'
export default {
  name: 'PodBrief',
  props: {
    pods: {
      type: Object,
      default: () => {}
    }
  },
  setup() {
    return {
      // filter
      statusPodFilter,
      // time format
      formatDate
    }
  }
}
</script>
