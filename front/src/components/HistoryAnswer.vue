<script setup>
export default {
  name: "HistoryAnswer",
  setup() {
    return {}
  }
}
</script>

<template>
  <el-select
    placeholder="选择历史记录">
    <el-option
      v-for="item in []"
      :key="item"
      :label="item"
      :value="item"
      @click=""
    />
  </el-select>
</template>
<style scoped>
</style>
