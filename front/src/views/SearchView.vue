<script>
import {useRoute, useRouter} from 'vue-router';
import {onMounted, ref} from 'vue';
import axios from "axios";

export default {
  name: "SearchView",
  setup() {
    const route = useRoute();
    const router = useRouter();
    const initialAnswer = route.query.answer || '';
    const answer = ref(initialAnswer);
    const initialPages = [];
    const pages = ref(initialPages);

    const fetchData = async () => {
      try {
        const url = `http://localhost:8080/page?text=${answer.value}`;
        const response = await axios.get(url);
        pages.value = response.data.data;
      } catch (err) {
        console.error(err);
      }
    };

    const handleSearch = () => {
      router.replace({query: {answer: answer.value}});
      fetchData();
    };

    onMounted(fetchData);

    return {
      answer,
      pages,
      handleSearch,
    };
  },
};
</script>

<template>
  <el-container>
    <el-header>
      <el-row style="margin-top:50px; margin-left: 20px">
        <img alt="logo" src="@/assets/logo.svg" style="width: 40px; height: 40px;"/>
        <el-input v-model="answer" style="width: 480px; height: 40px; margin-left: 10px"/>
        <el-button type="primary" style="height: 40px; margin-left: 5px" @click="handleSearch">
          <strong>Search</strong>
        </el-button>
      </el-row>
      <el-divider/>
    </el-header>
    <el-main style="margin-top:50px; margin-left: 50px">
      <div v-for="page in pages">
        <el-row>
          <a :href="page.url" style="margin-top: 15px" target="_blank">
            <text style="font-size: 20px">{{ page.title }}</text>
          </a>
        </el-row>
        <el-row>
          <text style="margin-top:10px; font-size: 16px">{{ page.describe }}</text>
        </el-row>
        <el-divider/>
      </div>
    </el-main>
    <el-footer>
    </el-footer>
  </el-container>
</template>
<style scoped>
</style>
