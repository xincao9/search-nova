<script>
import Logo from "@/assets/logo.svg";

export default {
  name: "SearchView",
  computed: {
    Logo() {
      return Logo
    }
  }
}
</script>

<template>
  <el-container>
    <el-header>
      <el-row style="margin-top:50px; margin-left: 20px">
        <el-image :src="Logo" style="width: 40px; height: 40px;"/>
        <el-input style="width: 480px; height: 40px; margin-left: 10px"/>
      </el-row>
    </el-header>
    <el-main>
      <el-row>

      </el-row>
    </el-main>
    <el-footer>
    </el-footer>
  </el-container>
</template>

<style scoped>

</style>
