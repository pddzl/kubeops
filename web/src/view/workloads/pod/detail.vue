<template>
  <div>
    <div class="detail-operation">
      <el-button icon="view" size="small" type="primary" plain @click="viewPod">查看</el-button>
      <el-button icon="expand" size="small" type="primary" plain @click="routerPod('log')">日志</el-button>
      <el-button icon="expand" size="small" type="primary" plain @click="routerPod('terminal')">终端</el-button>
      <el-button icon="delete" size="small" type="danger" plain @click="deleteFunc">删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="podDetail.metadata" title="元数据" name="metadata">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>名称</p>
                <span class="content">{{ podDetail.metadata.name }}</span>
              </div>
              <div class="item">
                <p>命名空间</p>
                <span class="content">{{ podDetail.metadata.namespace }}</span>
              </div>
              <div class="item">
                <p>UID</p>
                <span class="content">{{ podDetail.metadata.uid }}</span>
              </div>
              <div class="item">
                <p>创建时间</p>
                <span class="content">{{ formatDate(podDetail.metadata.creationTimestamp) }}</span>
              </div>
            </div>
            <div v-if="podDetail.metadata.labels" class="row">
              <div class="item">
                <p>标签</p>
                <div class="content">
                  <span v-for="(label, index) in podDetail.metadata.labels" :key="index" class="shadow">{{ index }}: {{
                    label
                  }}</span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.spec && podDetail.status" title="资源信息" name="resource">
          <div class="info-box">
            <div class="row">
              <div v-if="podDetail.spec.nodeName" class="item">
                <p>节点</p>
                <span class="content">{{ podDetail.spec.nodeName }}</span>
              </div>
              <div v-if="podDetail.status.phase" class="item">
                <p>状态</p>
                <el-tag :type="statusPodFilter(podDetail.status.phase)" size="small">
                  {{ podDetail.status.phase }}
                </el-tag>
              </div>
              <div v-if="podDetail.status.podIP" class="item">
                <p>IP</p>
                <span class="content">{{ podDetail.status.podIP }}</span>
              </div>
              <div v-if="podDetail.status.qosClass" class="item">
                <p>服务质量</p>
                <span class="content">{{ podDetail.status.qosClass }}</span>
              </div>
              <div v-if="podDetail.spec.restartPolicy" class="item">
                <p>重启策略</p>
                <span class="content">{{ podDetail.spec.restartPolicy }}</span>
              </div>
              <div v-if="podDetail.spec.serviceAccountName">
                <p>服务账号</p>
                <span class="content">{{ podDetail.spec.serviceAccountName }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.status?.conditions" title="状态" name="conditions">
          <div class="info-table">
            <el-table :data="podDetail.status.conditions">
              <el-table-column prop="type" label="类别" />
              <el-table-column prop="status" label="状态" />
              <el-table-column label="最后检查时间">
                <template #default="scope">
                  <span v-if="scope.row.lastProbeTime">{{ formatDate(scope.row.lastProbeTime) }}</span>
                  <span v-else>-</span>
                </template>
              </el-table-column>
              <el-table-column label="最后迁移时间">
                <template #default="scope">
                  <span v-if="scope.row.lastTransitionTime">{{ formatDate(scope.row.lastTransitionTime) }}</span>
                  <span v-else>-</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.metadata?.ownerReferences" title="控制器" name="controller">
          <div v-for="reference in podDetail.metadata.ownerReferences" :key="reference" class="info-box">
            <div class="row">
              <div class="item">
                <p>名称</p>
                <div class="content">
                  <router-link v-if="reference.kind === 'ReplicaSet'" :to="{ name: 'replicaSet_detail', query: { replicaSet: reference.name, namespace: namespace } }">
                    <el-link type="primary" :underline="false">{{ reference.name }}</el-link>
                  </router-link>
                  <router-link v-else-if="reference.kind === 'DaemonSet'" :to="{ name: 'daemonSet_detail', query: { daemonSet: reference.name, namespace: namespace } }">
                    <el-link type="primary" :underline="false">{{ reference.name }}</el-link>
                  </router-link>
                  <router-link v-else-if="reference.kind === 'Job'" :to="{ name: 'job_detail', query: { job: reference.name, namespace: namespace } }">
                    <el-link type="primary" :underline="false">{{ reference.name }}</el-link>
                  </router-link>
                </div>
              </div>
              <div class="item">
                <p>类型</p>
                <span class="content">{{ reference.kind }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.spec?.containers.length > 0" title="容器" name="container">
          <container :data="podDetail.spec.containers" />
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.spec?.initContainers && podDetail.spec.initContainers.length > 0" title="初始容器" name="initContainers">
          <container :data="podDetail.spec.initContainers" />
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="podFormat" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { statusPodFilter } from '@/mixin/filter.js'
import { getPodDetail, getPodRaw, deletePod } from '@/api/kubernetes/pod'
import { formatDate } from '@/utils/format'
import Container from './components/container.vue'
import VueCodeMirror from '@/components/codeMirror/index.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
export default {
  name: 'PodDetail',
  components: {
    Container,
    VueCodeMirror
  },
  setup() {
    // 折叠面板
    const activeNames = ref(['metadata', 'resource', 'conditions', 'controller', 'container', 'initContainers'])

    // 路由
    const route = useRoute()
    const pod = route.query.pod
    const namespace = route.query.namespace

    const router = useRouter()

    // 加载pod详情
    const podDetail = ref({})

    const getData = async() => {
      await getPodDetail({ pod: pod, namespace: namespace }).then(response => {
        if (response.code === 0) {
          podDetail.value = response.data
        }
      })
    }
    getData()

    // 查看编排
    const dialogFormVisible = ref(false)
    const podFormat = ref({})

    const viewPod = async() => {
      const result = await getPodRaw({ pod: pod, namespace: namespace })
      if (result.code === 0) {
        podFormat.value = JSON.stringify(result.data)
      }
      dialogFormVisible.value = true
    }

    const deleteFunc = async() => {
      ElMessageBox.confirm('此操作将永久删除该Pod, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deletePod({ namespace: namespace, pod: pod })
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '删除成功!'
            })
          }
        })
    }

    // 跳转日志、webssh页面
    const routerPod = async(dest) => {
      if (dest === 'log') {
        router.push({ name: 'pod_log', query: { pod: pod, namespace: namespace }})
      } else if (dest === 'terminal') {
        router.push({ name: 'pod_terminal', query: { pod: pod, namespace: namespace }})
      }
    }

    return {
      // 折叠面板
      activeNames,
      // data
      namespace,
      podDetail,
      // 查看编排
      podFormat,
      dialogFormVisible,
      viewPod,
      // time format
      formatDate,
      // filter
      statusPodFilter,
      // 操作
      routerPod,
      // 删除
      deleteFunc
    }
  }
}
</script>
