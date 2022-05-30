<template>
  <div ref="echart" class="pod" />
</template>

<script>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref } from 'vue'
import { getNodeType } from '@/api/kubernetes/node.js'
export default {
  name: 'PodStatus',
  setup() {
    const data = ref([])
    const getData = async() => {
      const res = await getNodeType()
      if (res.code === 0) {
        data.value = res.data
      }
    }

    const chart = ref(null)
    const echart = ref(null)
    const initChart = () => {
      chart.value = echarts.init(echart.value)
      setOptions()
    }

    const setOptions = () => {
      chart.value.setOption({
        title: {
          text: '节点类型',
          textStyle: {
            fontSize: 12
          },
          left: '35%'
        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b}: {c} ({d}%)'
        },
        color: ['#73c0de', '#FFA07A'],
        series: [
          {
            name: '状态',
            type: 'pie',
            radius: '60%',
            avoidLabelOverlap: false,
            label: {
              show: true,
              position: 'inside',
              formatter: '{d}%'
            },
            data: data.value,
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              },
            },
          },
        ],
      })
    }

    onMounted(async() => {
      await nextTick()
      await getData()
      initChart()
    })

    onUnmounted(() => {
      if (!chart.value) {
        return
      }
      chart.value.dispose()
      chart.value = null
    })

    return {
      echart
    }
  }
}
</script>

<style lang="scss" scoped>
.pod {
  height: 200px;
  width: 200px;
}
</style>
