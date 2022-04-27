<template>
  <div class="format">
    <!-- eslint-disable-next-line vue/attribute-hyphenation -->
    <vue-json-pretty :data="data" :deep="1" :showLine="true" />
  </div>
</template>

<script>
import { toRefs } from 'vue'
import VueJsonPretty from 'vue-json-pretty'
import 'vue-json-pretty/lib/styles.css'
export default {
  name: 'JsonPretty',
  components: {
    VueJsonPretty,
  },
  props: {
    data: {
      type: Object,
      default: () => {}
    },
    color: {
      type: String,
      default: () => 'dark'
    }
  },
  setp(props) {
    const { color } = toRefs(props)
    return {
      color
    }
  }
}
</script>

<style lang="scss" scoped>
.format {
  color: v-bind(color);
}
</style>
