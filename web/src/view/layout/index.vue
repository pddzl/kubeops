<template>
  <el-container class="layout-cont">
    <el-container :class="[isSider?'openside':'hideside',isMobile?'mobile':'']">
      <el-row :class="[isShadowBg?'shadowBg':'']" @click="changeShadow()" />
      <el-aside class="main-cont main-left">
        <div class="tilte" :style="{background: backgroundColor}">
          <img alt class="logoimg" src="@/assets/logo.png">
          <div v-if="isSider" class="tit-text" :style="{color:textColor}">KUBEOPS</div>
        </div>
        <Aside class="aside" />
      </el-aside>
    </el-container>
  </el-container>
</template>

<script>
export default {
  name: 'Layout'
}
</script>

<script setup>
import Aside from '@/view/layout/aside/index.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { emitter } from '@/utils/bus.js'

const router = useRouter()

// 三种窗口适配
const isCollapse = ref(false)
const isSider = ref(true)
const isMobile = ref(false)

const initPage = () => {
  const screenWidth = document.body.clientWidth
  if (screenWidth < 1000) {
    isMobile.value = true
    isSider.value = false
    isCollapse.value = true
  } else if (screenWidth >= 1000 && screenWidth < 1200) {
    isMobile.value = false
    isSider.value = false
    isCollapse.value = true
  } else {
    isMobile.value = false
    isSider.value = true
    isCollapse.value = false
  }
}
initPage()

const isShadowBg = ref(false)
const totalCollapse = () => {
  isCollapse.value = !isCollapse.value
  isSider.value = !isCollapse.value
  isShadowBg.value = !isCollapse.value
  emitter.emit('collapse', isCollapse.value)
}

const toPerson = () => {
  router.push({ name: 'person' })
}
const changeShadow = () => {
  isShadowBg.value = !isShadowBg.value
  isSider.value = !!isCollapse.value
  totalCollapse()
}

</script>

<style lang="scss">
@import '@/style/mobile.scss';

.dark{
  background-color: #191a23 !important;
  color: #fff !important;
}
.light{
  background-color: #fff !important;
  color: #000 !important;
}
</style>
