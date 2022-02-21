<template>
  <div class="row_container">
    <div v-for="(detail, index) in data" :key="index" class="container">
      <div>
        <el-tag size="small">{{ detail.name }}</el-tag>
      </div>
      <div class="row_mine">
        <p>镜像</p>
        {{ detail.image }}
      </div>
      <div v-if="detail.status" class="row_mine">
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
      <div v-if="detail.env && detail.env.length > 0" class="row_mine">
        <p>环境变量</p>
        <span
          v-for="(env, indexEnv) in detail.env"
          :key="indexEnv"
          class="span-shadow"
        >{{ env.name }}: {{ env.value }}</span>
      </div>
      <div v-if="detail.commands && detail.commands.length > 0" class="row_mine">
        <p>命令</p>
        <div class="div-shadow">
          <div v-for="(command, indexC) in detail.commands" :key="indexC">{{ command }}</div>
        </div>
      </div>
      <div v-if="detail.args && detail.args.length > 0" class="row_mine">
        <p>参数</p>
        <div class="div-shadow">
          <div v-for="(arg, indexA) in detail.args" :key="indexA">{{ arg }}</div>
        </div>
      </div>
      <div v-if="detail.volumeMounts && detail.volumeMounts.length > 0" class="row_mine">
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
      <div v-if="detail.securityContext" class="row_mine">
        <p>安全性上下文</p>
        <div class="row_context">
          <div
            v-if="detail.securityContext.capabilities.add && detail.securityContext.capabilities.add.length > 0"
          >
            <p>Added Capabilities</p>
            <span v-for="(add, indexA) in detail.securityContext.capabilities.add" :key="indexA">
              <span class="content">{{ add }}</span>
              <span v-if="detail.securityContext.capabilities.add.length - 1 != indexA">,</span>
            </span>
          </div>
          <div
            v-if="detail.securityContext.capabilities.drop && detail.securityContext.capabilities.drop.length > 0"
          >
            <p>Dropped Capabilities</p>
            <span v-for="(drop, indexD) in detail.securityContext.capabilities.drop" :key="indexD">
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
            <span class="content">{{ detail.securityContext.windowsOptions.gmsaCredentialSpecName }}</span>
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
      <div v-if="detail.livenessProbe" class="row_mine">
        <p>Liveness Probe</p>
        <div class="row_context">
          <div v-if="detail.livenessProbe.initialDelaySeconds">
            <p>Initial Delay (Seconds)</p>
            <span class="content">{{ detail.livenessProbe.initialDelaySeconds }}</span>
          </div>
          <div v-if="detail.livenessProbe.timeoutSeconds">
            <p>Timeout (Seconds)</p>
            <span class="content">{{ detail.livenessProbe.timeoutSeconds }}</span>
          </div>
          <div v-if="detail.livenessProbe.periodSeconds">
            <p>Probe Period (Seconds)</p>
            <span class="content">{{ detail.livenessProbe.periodSeconds }}</span>
          </div>
          <div v-if="detail.livenessProbe.successThreshold">
            <p>Success Threshold</p>
            <span class="content">{{ detail.livenessProbe.successThreshold }}</span>
          </div>
          <div v-if="detail.livenessProbe.failureThreshold">
            <p>Failure Threshold</p>
            <span class="content">{{ detail.livenessProbe.failureThreshold }}</span>
          </div>
          <div v-if="detail.livenessProbe.httpGet">
            <p>HTTP Healthcheck URI</p>
            <span
              class="content_shadow"
            >{{ detail.livenessProbe.httpGet.scheme + '://[host]:' + detail.livenessProbe.httpGet.port + detail.livenessProbe.httpGet.path }}</span>
          </div>
        </div>
      </div>
      <div v-if="detail.readinessProbe" class="row_mine">
        <p>ReadinessProbe</p>
        <div class="row_context">
          <div v-if="detail.readinessProbe.initialDelaySeconds">
            <p>Initial Delay (Seconds)</p>
            <span class="content">{{ detail.readinessProbe.initialDelaySeconds }}</span>
          </div>
          <div v-if="detail.readinessProbe.timeoutSeconds">
            <p>Timeout (Seconds)</p>
            <span class="content">{{ detail.readinessProbe.timeoutSeconds }}</span>
          </div>
          <div v-if="detail.readinessProbe.periodSeconds">
            <p>Probe Period (Seconds)</p>
            <span class="content">{{ detail.readinessProbe.periodSeconds }}</span>
          </div>
          <div v-if="detail.readinessProbe.successThreshold">
            <p>Success Threshold</p>
            <span class="content">{{ detail.readinessProbe.successThreshold }}</span>
          </div>
          <div v-if="detail.readinessProbe.failureThreshold">
            <p>Failure Threshold</p>
            <span class="content">{{ detail.readinessProbe.failureThreshold }}</span>
          </div>
          <div v-if="detail.readinessProbe.httpGet">
            <p>HTTP Healthcheck URI</p>
            <span
              class="content_shadow"
            >{{ detail.readinessProbe.httpGet.scheme + '://[host]:' + detail.livenessProbe.httpGet.port + detail.livenessProbe.httpGet.path }}</span>
          </div>
        </div>
      </div>
      <div v-if="detail.startupProbe" class="row_mine">
        <p>StartupProbe</p>
        <div class="row_context">
          <div v-if="detail.startupProbe.initialDelaySeconds">
            <p>Initial Delay (Seconds)</p>
            <span class="content">{{ detail.startupProbe.initialDelaySeconds }}</span>
          </div>
          <div v-if="detail.startupProbe.timeoutSeconds">
            <p>Timeout (Seconds)</p>
            <span class="content">{{ detail.startupProbe.timeoutSeconds }}</span>
          </div>
          <div v-if="detail.startupProbe.periodSeconds">
            <p>Probe Period (Seconds)</p>
            <span class="content">{{ detail.startupProbe.periodSeconds }}</span>
          </div>
          <div v-if="detail.startupProbe.successThreshold">
            <p>Success Threshold</p>
            <span class="content">{{ detail.startupProbe.successThreshold }}</span>
          </div>
          <div v-if="detail.startupProbe.failureThreshold">
            <p>Failure Threshold</p>
            <span class="content">{{ detail.startupProbe.failureThreshold }}</span>
          </div>
          <div v-if="detail.startupProbe.httpGet">
            <p>HTTP Healthcheck URI</p>
            <span
              class="content_shadow"
            >{{ detail.startupProbe.httpGet.scheme + '://[host]:' + detail.livenessProbe.httpGet.port + detail.livenessProbe.httpGet.path }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { formatDate } from '@/utils/format'
export default {
  name: 'Container',
  props: {
    data: {
      type: Object,
      required: true
    }
  },
  setup() {
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
      volumeFilter,
      formatDate
    }
  }
}
</script>

<style lang="scss" scoped>
.row_container {
  > div:not(:last-child) {
    margin-bottom: 40px;
  }
  .container {
    > div:not(:last-child) {
      margin-bottom: 20px;
    }
  }
  .table {
    margin-right: 20px;
  }
}
.content_shadow {
  background-color: rgba(128, 128, 128, 0.253);
  padding: 5px 15px 5px 15px;
  border-radius: 5px;
  display: inline-block;
  margin-top: 5px;
  box-sizing: border-box;
}
</style>
