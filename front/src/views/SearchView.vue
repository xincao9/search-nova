<script>
import logo from "@/assets/img/logo_big.png";

export default {
  name: "SearchView",
  computed: {
    logo() {
      return logo
    }
  }
}
</script>

<template>
  <el-container>
    <el-header>
      <el-row style="margin-top:50px; margin-left: 20px">
        <el-image :src="logo" style="width: 40px; height: 40px;"/>
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
