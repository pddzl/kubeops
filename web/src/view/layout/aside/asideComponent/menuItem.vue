<template>
  <el-menu-item :index="routerInfo.name" :route="{parameters:routerInfo.parameters}">
    <el-icon>
      <component :is="routerInfo.meta.icon" />
    </el-icon>
    <template #title>
      <span class="gva-menu-item-title">{{ routerInfo.meta.title }}</span>
    </template>
  </el-menu-item>
</template>

<script>
export default {
  name: 'MenuItem',
}
</script>

<script setup>
const props = defineProps({
  routerInfo: {
    default: function() {
      return null
    },
    type: Object
  }
})
console.log('parameters', props.routerInfo.parameters)
</script>

<style lang="scss" scoped>
// .gva-menu-item-title {
//   min-width: 160px;
// }
</style>
