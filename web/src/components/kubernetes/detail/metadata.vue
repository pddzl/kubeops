<template>
  <div class="info-box">
    <div class="row">
      <div v-if="metadata.name" class="item">
        <p>名称</p>
        <span class="content">{{ metadata.name }}</span>
      </div>
      <div v-if="metadata.namespace" class="item">
        <p>命名空间</p>
        <span class="content">{{ metadata.namespace }}</span>
      </div>
      <div v-if="metadata.uid" class="item">
        <p>UID</p>
        <span class="content">{{ metadata.uid }}</span>
      </div>
      <div v-if="metadata.creationTimestamp">
        <p>创建时间</p>
        <span class="content">{{ formatDate(metadata.creationTimestamp) }}</span>
      </div>
    </div>
    <div v-if="metadata.labels" class="row">
      <div class="item">
        <p>标签:</p>
        <div class="content">
          <span v-for="(label, index) in metadata.labels" :key="index" class="shadow">
            {{ index }}
            <span v-if="label">:</span>
            {{ label }}
          </span>
        </div>
      </div>
    </div>
    <div v-if="JSON.stringify(annotationsFormat) !== '{}'">
      <p style="font-size: 13px;">注释:</p>
      <vue-json-pretty :data="annotationsFormat" :color="'lightcoral'" />
    </div>
  </div>
</template>

<script>
import { ref, toRefs } from 'vue'
import { formatDate } from '@/utils/format'
import VueJsonPretty from '@/components/vueJsonPretty/index.vue'
export default {
  name: 'MetaData',
  components: {
    VueJsonPretty
  },
  props: {
    metadata: {
      type: Object,
      default: () => { }
    }
  },
  setup(props) {
    const { metadata } = toRefs(props)
    const annotationsFormat = ref({})

    // 去掉双引号、反斜杠、换行符
    const format = () => {
      if (metadata.value.annotations) {
        try {
          annotationsFormat.value = JSON.parse(JSON.stringify(metadata.value.annotations).replace(/\\"/g, '"').replace(/"\{/g, '{').replace(/\\n"/g, ''))
        } catch (err) {
          annotationsFormat.value = metadata.value.annotations
        }
      }
    }
    format()

    return {
      // 响应式数据
      annotationsFormat,
      formatDate
    }
  }
}
</script>
