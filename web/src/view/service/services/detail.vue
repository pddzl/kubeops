<template>
  <div>
    <div class="detail-operation">
      <div class="button">
        <el-affix :offset="120">
          <el-button icon="view" size="small" type="primary" plain @click="editServicesSet">查看</el-button>
          <el-button icon="delete" size="small" type="danger" plain>删除</el-button>
        </el-affix>
      </div>
    </div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="servicesDetail.metadata" title="元数据" name="metadata">
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>名称</p>
              <span class="content">{{ servicesDetail.metadata.name }}</span>
            </div>
            <div>
              <p>命名空间</p>
              <span class="content">{{ servicesDetail.metadata.namespace }}</span>
            </div>
            <div>
              <p>UID</p>
              <span class="content">{{ servicesDetail.metadata.uid }}</span>
            </div>
            <div>
              <p>创建时间</p>
              <span class="content">{{ formatDate(servicesDetail.metadata.creationTimestamp) }}</span>
            </div>
          </div>
        </div>
        <div class="common_show">
          <p>标签:</p>
          <span v-for="(label, index) in servicesDetail.metadata.labels" :key="index" class="span-shadow">
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
      <el-collapse-item v-if="servicesDetail.spec" title="资源信息" name="spec">
        <div class="row_mine" style="margin-bottom: 10px;">
          <div class="row_context">
            <div>
              <p>类别</p>
              <span class="content">{{ servicesDetail.spec.type }}</span>
            </div>
            <div>
              <p>集群IP</p>
              <span class="content">{{ servicesDetail.spec.clusterIP }}</span>
            </div>
            <div>
              <p>sessionAffinity</p>
              <span class="content">{{ servicesDetail.spec.sessionAffinity }}</span>
            </div>
          </div>
        </div>
        <div class="row_mine">
          <div class="row_context">
            <div>
              <p>Selector</p>
              <span v-for="(label, index) in servicesDetail.spec.selector" :key="index" class="span-shadow">
                {{ index }}
                <span v-if="label">:</span>
                {{ label }}
              </span>
            </div>
          </div>
        </div>
      </el-collapse-item>
      <el-collapse-item title="Pods" name="pods">
        <div style="padding-right: 20px;">
          <el-table :data="servicesPods">
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
    </el-collapse>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="servicesFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { getServicesDetail, getServicesRaw, getServicesPods } from '@/api/kubernetes/services'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'
import { formatDate } from '@/utils/format'
import { statusRsFilter } from '@/mixin/filter.js'
import { statusPodFilter } from '@/mixin/filter.js'
export default {
  name: 'ServicesDetail',
  components: {
    VueCodeMirror,
    VueJsonPretty
  },
  setup() {
    const activeNames = ref(['metadata', 'spec', 'pods'])
    const page = ref(1)
    const pageSize = ref(10)
    const total = ref(0)
    const servicesPods = ref([])
    const servicesDetail = ref({})
    const servicesFormat = ref({})
    const annotationsFormat = ref({})
    const dialogFormVisible = ref(false)

    const route = useRoute()
    const namespace = route.query.namespace
    const service = route.query.service

    // 获取daemonSet详情
    const getServicesDetailData = async() => {
      await getServicesDetail({ namespace: namespace, service: service }).then(response => {
        if (response.code === 0) {
          servicesDetail.value = response.data
          // annotationsFormat.value = JSON.stringify(response.data.metadata.annotations, null, 2)
          annotationsFormat.value = JSON.parse(JSON.stringify(response.data.metadata.annotations).replace(/\\"/g, '"').replace(/"\{/g, '{').replace(/\\n"/g, ''))
        }
      })
    }
    getServicesDetailData()

    // 加载关联pods
    const getServicesPodsData = async() => {
      const table = await getServicesPods({ page: page.value, pageSize: pageSize.value, namespace: namespace, service: service })
      if (table.code === 0) {
        servicesPods.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
      }
    }
    getServicesPodsData()

    // 操作
    const editServicesSet = async() => {
      const result = await getServicesRaw({ service: service, namespace: namespace })
      if (result.code === 0) {
        servicesFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    // 分页
    const handleSizeChange = (val) => {
      pageSize.value = val
      getServicesPodsData()
    }

    const handleCurrentChange = (val) => {
      page.value = val
      getServicesPodsData()
    }

    return {
      // 响应式数据
      activeNames,
      servicesDetail,
      servicesFormat,
      annotationsFormat,
      dialogFormVisible,
      servicesPods,
      // filter
      statusRsFilter,
      statusPodFilter,
      // 格式化日期
      formatDate,
      // 操作
      editServicesSet,
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
