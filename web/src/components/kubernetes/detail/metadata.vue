<template>
  <div>
    <div class="row_mine">
      <div class="row_context">
        <div v-if="metadata.name">
          <p>名称</p>
          <span class="content">{{ metadata.name }}</span>
        </div>
        <div v-if="metadata.namespace">
          <p>命名空间</p>
          <span class="content">{{ metadata.namespace }}</span>
        </div>
        <div v-if="metadata.uid">
          <p>UID</p>
          <span class="content">{{ metadata.uid }}</span>
        </div>
        <div v-if="metadata.creationTimestamp">
          <p>创建时间</p>
          <span class="content">{{ formatDate(metadata.creationTimestamp) }}</span>
        </div>
      </div>
    </div>
    <div v-if="metadata.labels" class="metadata-label">
      <p>标签:</p>
      <div class="label-flex">
        <div v-for="(label, index) in metadata.labels" :key="index" class="label">
          <span>
            {{ index }}
            <span v-if="label">:</span>
            {{ label }}
          </span>
        </div>
      </div>
    </div>
    <div v-if="JSON.stringify(annotationsFormat) !== '{}'" style="margin-top: 10px;">
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
        annotationsFormat.value = JSON.parse(JSON.stringify(metadata.value.annotations).replace(/\\"/g, '"').replace(/"\{/g, '{').replace(/\\n"/g, ''))
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

<style lang="scss" scoped>
.metadata-label {
  margin-top: 10px;
  p {
    margin-bottom: 5px;
    font-size: 13px;
  }
  .label-flex {
    display: flex;
    flex-wrap: wrap;
    div:not(:last-child) {
      margin-right: 8px;
    }
    div {
      padding: 2px;
      >span {
        background-color: rgba(128, 128, 128, 0.159);
        font-size: 13px;
        border-radius: 8px;
        padding: 5px;
      }
    }
  }
}
</style>
