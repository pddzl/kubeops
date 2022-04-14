<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="editDaemonSet">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="DaemonSetDetail.metadata" title="元数据" name="1">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>名称</p>
              <span class="content">{{ DaemonSetDetail.metadata.name }}</span>
            </div>
            <div>
              <p>命名空间</p>
              <span class="content">{{ DaemonSetDetail.metadata.namespace }}</span>
            </div>
            <div>
              <p>UID</p>
              <span class="content">{{ DaemonSetDetail.metadata.uid }}</span>
            </div>
            <div>
              <p>创建时间</p>
              <span class="content">{{ formatDate(DaemonSetDetail.metadata.creationTimestamp) }}</span>
            </div>
          </div>
        </div>
        <div class="common_show">
          <p>标签:</p>
          <span v-for="(label, index) in DaemonSetDetail.metadata.labels" :key="index" class="span-shadow">
            {{ index }}
            <span v-if="label">:</span>
            {{ label }}
          </span>
        </div>
        <div v-if="annotationsFormat" class="common_show">
          <p>注释:</p>
          <div style="color: lightcoral;">
            <!-- eslint-disable-next-line vue/attribute-hyphenation -->
            <vue-json-pretty :data="annotationsFormat" :deep="2" :showLine="true" />
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="DaemonSetDetail.spec" title="资源信息" name="2">
        <div class="row_mine" style="margin-bottom: 10px;">
          <div class="row_context">
            <div>
              <p>Selector</p>
              <span v-for="(label, index) in DaemonSetDetail.spec.selector" :key="index" class="span-shadow">
                {{ index }}
                <span v-if="label">:</span>
                {{ label }}
              </span>
            </div>
          </div>
        </div>
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>Strategy</p>
              <span class="content">{{ DaemonSetDetail.spec.updateStrategy.type }}</span>
            </div>
            <div v-if="DaemonSetDetail.spec.updateStrategy.type === 'RollingUpdate'">
              <p>maxUnavailable</p>
              <span class="content">{{ DaemonSetDetail.spec.updateStrategy.rollingUpdate.maxUnavailable }}</span>
            </div>
            <div v-if="DaemonSetDetail.spec.updateStrategy.type === 'RollingUpdate'">
              <p>maxSurge</p>
              <span class="content">{{ DaemonSetDetail.spec.updateStrategy.rollingUpdate.maxUnavailable }}</span>
            </div>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="DaemonSetDetail.status" title="状态" name="3">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>currentNumberScheduled</p>
              <span class="content">{{ DaemonSetDetail.status.currentNumberScheduled }}</span>
            </div>
            <div>
              <p>desiredNumberScheduled</p>
              <span class="content">{{ DaemonSetDetail.status.desiredNumberScheduled }}</span>
            </div>
            <div>
              <p>numberReady</p>
              <span class="content">{{ DaemonSetDetail.status.numberReady }}</span>
            </div>
            <div>
              <p>numberAvailable</p>
              <span class="content">{{ DaemonSetDetail.status.numberAvailable }}</span>
            </div>
          </div>
        </div>
        <div v-if="DaemonSetDetail.status.conditions" style="margin-top: 20px; margin-right: 20px;">
          <el-table :data="DaemonSetDetail.status.conditions">
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
    </el-collapse>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="daemonSetFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getDaemonSetDetail, getDaemonSetRaw } from '@/api/kubernetes/daemonSet'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'
import { formatDate } from '@/utils/format'
import { statusRsFilter } from '@/mixin/filter.js'
export default {
  name: 'DaemonSetDetail',
  components: {
    VueCodeMirror,
    VueJsonPretty
  },
  setup() {
    const activeNames = ref(['1', '2', '3'])
    const DaemonSetDetail = ref({})
    const daemonSetFormat = ref({})
    const annotationsFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const daemonSet = route.query.daemonSet

    // 获取daemonSet详情
    const getDaemonSetDetailData = async() => {
      await getDaemonSetDetail({ namespace: namespace, daemonSet: daemonSet }).then(response => {
        if (response.code === 0) {
          DaemonSetDetail.value = response.data
          // annotationsFormat.value = JSON.stringify(response.data.metadata.annotations, null, 2)
          annotationsFormat.value = JSON.parse(JSON.stringify(response.data.metadata.annotations).replace(/\\"/g, '"').replace(/"\{/g, '{').replace(/\\n"/g, ''))
        }
      })
    }
    getDaemonSetDetailData()

    // 操作
    const editDaemonSet = async() => {
      const result = await getDaemonSetRaw({ daemonSet: daemonSet, namespace: namespace })
      if (result.code === 0) {
        daemonSetFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    return {
      // 响应式数据
      activeNames,
      DaemonSetDetail,
      daemonSetFormat,
      annotationsFormat,
      dialogFormVisible,
      // status filter
      statusRsFilter,
      // 格式化日期
      formatDate,
      // 操作
      editDaemonSet
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
