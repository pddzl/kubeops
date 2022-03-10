<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true">
        <el-form-item label="容器">
          <el-select v-model="searchInfo.namespace" clearable placeholder="请选择">
            <el-option v-for="item in namespace" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="行数">
          <el-select v-model="searchInfo.namespace" clearable placeholder="请选择">
            <el-option v-for="item in namespace" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="log">
      <p>here</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Log',
  props: {
    logRaw: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="scss" scoped>
.log {
  padding: 24px;
  background-color: rgba(0, 0, 0, 0.178);
  border-radius: 2px;
}
</style>
