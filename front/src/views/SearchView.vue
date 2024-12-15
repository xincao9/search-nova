<script>
import Logo from "@/assets/logo.svg";
import {useRoute} from 'vue-router';

export default {
  name: "SearchView",
  setup() {
    const route = useRoute();
    const answer = route.query.answer;
    return {Logo, answer}
  }
}
</script>

<template>
  <el-container>
    <el-header>
      <el-row style="margin-top:50px; margin-left: 20px">
        <el-image :src="Logo" style="width: 40px; height: 40px;"/>
        <el-input style="width: 480px; height: 40px; margin-left: 10px" v-model="answer"/>
      </el-row>
    </el-header>
    <el-main>
    </el-main>
    <el-footer>
    </el-footer>
  </el-container>
</template>

<style scoped>

</style>
