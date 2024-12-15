<script setup>
import {Search} from '@element-plus/icons-vue'
import logo from '@/assets/img/logo_big.png'
</script>

<template>
  <el-container>
    <el-header>
    </el-header>
    <el-main>
      <el-row justify="center">
        <el-col :span="4">
          <el-image :src="logo"/>
        </el-col>
      </el-row>
      <el-row justify="center">
        <div style="margin-top: 80px">
          <el-input style="width: 480px; height: 40px;"/>
        </div>
      </el-row>
      <el-row justify="center">
        <div style="margin-top: 30px">
          <el-button plain size="large" :icon="Search">Good Search</el-button>
          <el-button plain size="large">I'm Feeling Lucky</el-button>
        </div>
      </el-row>
    </el-main>
    <el-footer>
    </el-footer>
  </el-container>
</template>

<style scoped>
</style>
