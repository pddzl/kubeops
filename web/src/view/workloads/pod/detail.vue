<template>
  <div>
    <el-collapse v-model="activeNames">
      <el-collapse-item v-if="podDetail.metadata" class="object_meta" title="元数据" name="1">
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
        <div class="row_mine">
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
      <el-collapse-item v-if="podDetail.containers" title="容器" name="5" class="container">
        <div v-for="(detail, index) in podDetail.containers" :key="index">
          <div class="bottom">
            <p>{{ detail.name }}</p>
          </div>
          <div class="bottom row_mine">
            <p>镜像</p>
            {{ detail.image }}
          </div>
          <div class="row_mine status">
            <p>状态</p>
            <div class="row_context">
              <div>
                <p>Ready</p>
                {{ detail.status.ready }}
              </div>
              <div>
                <p>Started</p>
                {{ detail.status.started }}
              </div>
              <div>
                <p>创建时间</p>
                {{ formatDate(detail.status.state.running.startedAt) }}
              </div>
            </div>
          </div>
          <div v-if="detail.env.length > 0" class="row_mine bottom">
            <p>环境变量</p>
            <span
              v-for="(env, indexEnv) in detail.env"
              :key="indexEnv"
              class="span-shadow"
            >{{ env.name }}: {{ env.value }}</span>
          </div>
          <div v-if="detail.commands" class="row_mine bottom">
            <p>命令</p>
            <div class="div-shadow">
              <div v-for="(command, indexC) in detail.commands" :key="indexC">{{ command }}</div>
            </div>
          </div>
          <div v-if="detail.args" class="row_mine bottom">
            <p>参数</p>
            <div class="div-shadow">
              <div v-for="(arg, indexA) in detail.args" :key="indexA">{{ arg }}</div>
            </div>
          </div>
          <div class="row_mine bottom">
            <p>挂载点</p>
            <div class="table">
              <el-table :data="detail.volumeMounts">
                <el-table-column prop="name" label="名称" min-width="120" />
                <el-table-column prop="readOnly" label="只读" min-width="80" />
                <el-table-column prop="mountPath" label="挂载路径" min-width="240" />
                <el-table-column label="子路径">
                  <template #default="scope">
                    <span v-if="scope.row.subPath">{{ scope.row.subPath }}</span>
                    <span v-else>-</span>
                  </template>
                </el-table-column>
                <el-table-column label="源类型">
                  <template #default="scope">{{ volumeFilter(scope.row.volume) }}</template>
                </el-table-column>
                <el-table-column label="源名称" min-width="120">
                  <template #default="scope">{{ volumeFilter(scope.row.volume, true) }}</template>
                </el-table-column>
              </el-table>
            </div>
          </div>
          <div class="row_mine">
            <p>安全性上下文</p>
            <div class="row_context">
              <div
                v-if="detail.securityContext.capabilities.add && detail.securityContext.capabilities.add.length > 0"
              >
                <p>Added Capabilities</p>
                <span
                  v-for="(add, indexA) in detail.securityContext.capabilities.add"
                  :key="indexA"
                >
                  <span class="content">{{ add }}</span>
                  <span v-if="detail.securityContext.capabilities.add.length - 1 != indexA">,</span>
                </span>
              </div>
              <div
                v-if="detail.securityContext.capabilities.drop && detail.securityContext.capabilities.drop.length > 0"
              >
                <p>Dropped Capabilities</p>
                <span
                  v-for="(drop, indexD) in detail.securityContext.capabilities.drop"
                  :key="indexD"
                >
                  <span class="content">{{ drop }}</span>
                  <span v-if="detail.securityContext.capabilities.drop.length - 1 != indexD">,</span>
                </span>
              </div>
              <div v-if="detail.securityContext.privileged !== undefined">
                <p>Privileged</p>
                <span class="content">{{ detail.securityContext.privileged }}</span>
              </div>
              <div v-if="detail.securityContext.seLinuxOptions?.user">
                <p>SeLinuxOptions User</p>
                <span class="content">{{ detail.securityContext.seLinuxOptions.user }}</span>
              </div>
              <div v-if="detail.securityContext.seLinuxOptions?.role">
                <p>SeLinuxOptions Role</p>
                <span class="content">{{ detail.securityContext.seLinuxOptions.role }}</span>
              </div>
              <div v-if="detail.securityContext.seLinuxOptions?.type">
                <p>SeLinuxOptions Type</p>
                <span class="content">{{ detail.securityContext.seLinuxOptions.type }}</span>
              </div>
              <div v-if="detail.securityContext.seLinuxOptions?.level">
                <p>SeLinuxOptions Level</p>
                <span class="content">{{ detail.securityContext.seLinuxOptions.level }}</span>
              </div>
              <div v-if="detail.securityContext.windowsOptions?.gmsaCredentialSpecName">
                <p>WindowsOptions GMSACredentialSpecName</p>
                <span
                  class="content"
                >{{ detail.securityContext.windowsOptions.gmsaCredentialSpecName }}</span>
              </div>
              <div v-if="detail.securityContext.windowsOptions?.gmsaCredentialSpec">
                <p>WindowsOptions GMSACredentialSpec</p>
                <span class="content">{{ detail.securityContext.windowsOptions.gmsaCredentialSpec }}</span>
              </div>
              <div v-if="detail.securityContext.windowsOptions?.runAsUserName">
                <p>WindowsOptions RunAsUserName</p>
                <span class="content">{{ detail.securityContext.windowsOptions.runAsUserName }}</span>
              </div>
              <div v-if="detail.securityContext.windowsOptions?.hostProcess != undefined">
                <p>WindowsOptions HostProcess</p>
                <span class="conetnt">{{ detail.securityContext.windowsOptions.hostProcess }}</span>
              </div>
              <div v-if="detail.securityContext.runAsUser">
                <p>RunAsUser</p>
                {{ detail.securityContext.runAsUser }}
              </div>
              <div v-if="detail.securityContext.runAsGroup">
                <p>RunAsGroup</p>
                <span class="content">{{ detail.securityContext.runAsGroup }}</span>
              </div>
              <div v-if="detail.securityContext.runAsNonRoot !== undefined">
                <p>RunAsNonRoot</p>
                <span class="content">{{ detail.securityContext.runAsNonRoot }}</span>
              </div>
              <div v-if="detail.securityContext.readOnlyRootFilesystem !== undefined">
                <p>ReadOnlyRootFilesystem</p>
                <span class="conetnt">{{ detail.securityContext.readOnlyRootFilesystem }}</span>
              </div>
              <div v-if="detail.securityContext.allowPrivilegeEscalation !== undefined">
                <p>AllowPrivilegeEscalation</p>
                <span class="content">{{ detail.securityContext.allowPrivilegeEscalation }}</span>
              </div>
              <div v-if="detail.securityContext.procMount">
                <p>ProcMount</p>
                <span class="content">{{ detail.securityContext.procMount }}</span>
              </div>
              <div v-if="detail.securityContext.seccompProfile?.type">
                <p>SeccompProfile Type</p>
                <span class="content">{{ detail.securityContext.seccompProfile.type }}</span>
              </div>
              <div v-if="detail.securityContext.seccompProfile?.localhostProfile">
                <p>SeccompProfile LocalhostProfile</p>
                <span class="content">{{ detail.securityContext.seccompProfile.localhostProfile }}</span>
              </div>
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
import { getPodDetail } from '@/api/kubernetes/pod'
import { formatDate } from '@/utils/format'
export default {
  name: 'PodDetail',
  setup() {
    const activeNames = ref(['1', '2', '3', '4', '5'])
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

    // filter
    const volumeFilter = (volume, source = false) => {
      if (Object.prototype.hasOwnProperty.call(volume, 'hostPath')) {
        if (source) {
          return volume.hostPath.path
        } else {
          return 'hostPath'
        }
      } else if (Object.prototype.hasOwnProperty.call(volume, 'configMap')) {
        if (source) {
          return volume.configMap.name
        } else {
          return 'configMap'
        }
      } else if (Object.prototype.hasOwnProperty.call(volume, 'projected')) {
        if (source) {
          return '-'
        } else {
          return 'projected'
        }
      } else {
        return '-'
      }
    }

    return {
      activeNames,
      podDetail,
      formatDate,
      volumeFilter
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
.container {
  .bottom {
    margin-bottom: 30px;
  }
  .status {
    margin-bottom: 20px;
  }
}
.table {
  margin-right: 20px;
}
.row_mine {
  > p:first-child {
    margin-bottom: 10px;
  }
  .row_context {
    display: flex;
    div:not(:last-child) {
      margin-right: 100px;
    }
    p {
      font-size: 13px;
    }
    .content {
      font-size: 16px;
    }
  }
}
</style>
