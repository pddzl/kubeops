<template>
  <div class="page">
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>集群</span>
        </div>
        <div style="margin-top: 20px;">
          <cluster-status />
        </div>
      </div>
    </div>
    <div class="gva-card-box">
      <el-card class="gva-card quick-entrance">
        <template #header>
          <div class="card-header">
            <span>快捷入口</span>
          </div>
        </template>
        <el-row :gutter="20">
          <el-col
            v-for="(card, key) in toolCards"
            :key="key"
            :span="4"
            :xs="8"
            class="quick-entrance-items"
            @click="toTarget(card.name)"
          >
            <div class="quick-entrance-item">
              <div class="quick-entrance-item-icon" :style="{ backgroundColor: card.bg }">
                <el-icon>
                  <component :is="card.icon" :style="{ color: card.color }" />
                </el-icon>
              </div>
              <p>{{ card.label }}</p>
            </div>
          </el-col>
        </el-row>
      </el-card>
    </div>
  </div>
</template>

<script setup>
// import echartsLine from '@/view/dashboard/dashboardCharts/echartsLine.vue'
// import dashboardTable from '@/view/dashboard/dashboardTable/dashboardTable.vue'
import ClusterStatus from '@/view/dashboard/cluster/index.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const toolCards = ref([
  {
    label: '节点',
    icon: 'monitor',
    name: 'node',
    color: '#ff9c6e',
    bg: 'rgba(255, 156, 110,.3)'
  },
  {
    label: '命名空间',
    icon: 'box',
    name: 'namespace',
    color: '#69c0ff',
    bg: 'rgba(105, 192, 255,.3)'
  },
  {
    label: '服务',
    icon: 'iphone',
    name: 'services',
    color: '#b37feb',
    bg: 'rgba(179, 127, 235,.3)'
  },
  {
    label: 'Deployment',
    icon: 'van',
    name: 'deployment',
    color: '#ffd666',
    bg: 'rgba(255, 214, 102,.3)'
  },
  {
    label: 'DaemonSet',
    icon: 'rank',
    name: 'daemonSet',
    color: '#ff85c0',
    bg: 'rgba(255, 133, 192,.3)'
  },
  {
    label: 'Pod',
    icon: 'orange',
    name: 'pod',
    color: '#5cdbd3',
    bg: 'rgba(92, 219, 211,.3)'
  },
])

const router = useRouter()

const toTarget = (name) => {
  router.push({ name })
}

</script>
<script>
export default {
  name: 'Dashboard'
}
</script>

<style lang="scss" scoped>
@mixin flex-center {
    display: flex;
    align-items: center;
}
.page {
    background: #f0f2f5;
    padding: 0;
    .gva-card-box{
      padding: 12px 16px;
      &+.gva-card-box{
        padding-top: 0px;
      }
    }
    .gva-card {
      box-sizing: border-box;
        background-color: #fff;
        border-radius: 2px;
        height: auto;
        padding: 26px 30px;
        overflow: hidden;
        box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
    }
     ::v-deep(.el-card__header){
          padding:0;
          border-bottom: none;
        }
        .card-header{
          padding-bottom: 20px;
          border-bottom: 1px solid #e8e8e8;
        }
    .quick-entrance-items {
        @include flex-center;
        justify-content: center;
        text-align: center;
        color: #333;
        .quick-entrance-item {
          padding: 16px 28px;
          margin-top: -16px;
          margin-bottom: -16px;
          border-radius: 4px;
          transition: all 0.2s;
          &:hover{
            box-shadow: 0px 0px 7px 0px rgba(217, 217, 217, 0.55);
          }
            cursor: pointer;
            height: auto;
            text-align: center;
            // align-items: center;
            &-icon {
                width: 50px;
                height: 50px !important;
                border-radius: 8px;
                @include flex-center;
                justify-content: center;
                margin: 0 auto;
                i {
                    font-size: 24px;
                }
            }
            p {
                margin-top: 10px;
            }
        }
    }
}
//小屏幕不显示右侧，将登陆框居中
@media (max-width: 750px) {
    .gva-card {
        padding: 20px 10px !important;
    }
}
</style>
