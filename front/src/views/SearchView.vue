<script>
import Logo from "@/assets/logo.svg";
import {useRoute} from 'vue-router';

export default {
  name: "SearchView",
  setup() {
    const route = useRoute();
    const answer = route.query.answer;
    const pages = [
      {
        "title": "澳门“一国两制”成功实践获得4点重要经验",
        "describe": "在热烈的掌声中，习近平发表了重要讲话。习近平指出，澳门地方虽小，但在“一国两制”实践中作用独特。总结澳门“一国两制”成功实践，可以获得以下4点重要经验",
        "url": "https://element-plus.org",
      },
      {
        "title": "澳门“一国两制”成功实践获得4点重要经验",
        "describe": "在热烈的掌声中，习近平发表了重要讲话。习近平指出，澳门地方虽小，但在“一国两制”实践中作用独特。总结澳门“一国两制”成功实践，可以获得以下4点重要经验",
        "url": "https://element-plus.org",
      }
    ]
    return {Logo, answer, pages}
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
      <el-row>
        <el-divider/>
      </el-row>
    </el-header>
    <el-main style="margin-top:50px; margin-left: 20px">
      <div v-for="page in pages">
        <el-row style="margin-top: 15px">
          <el-link href="{{ page.url }}" target="_blank">
            <el-text size="large">{{ page.title }}</el-text>
          </el-link>
        </el-row>
        <el-row style="margin-top:20px">
          <el-text size="default">{{ page.describe }}</el-text>
        </el-row>
      </div>
    </el-main>
    <el-footer>
    </el-footer>
  </el-container>
</template>

<style scoped>

</style>
