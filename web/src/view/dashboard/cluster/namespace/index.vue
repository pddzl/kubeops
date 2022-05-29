<template>
  <div ref="echart" class="namespace" />
</template>

<script>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref } from 'vue'
export default {
  name: 'NamespaceStatus',
  setup() {
    const data = [
      { value: 10, name: '正常' },
      { value: 0, name: '异常' }
    ]
    const chart = ref(null)
    const echart = ref(null)
    const initChart = () => {
      chart.value = echarts.init(echart.value)
      setOptions()
    }

    const setOptions = () => {
      chart.value.setOption({
        title: {
          text: '命名空间状态',
          textStyle: {
            fontSize: 12
          },
          left: '30%'
        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b}: {c} ({d}%)'
        },
        color: ['#91cc75', '#ee6666'],
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
            data: data,
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
.namespace {
  height: 200px;
  width: 200px;
}
</style>
