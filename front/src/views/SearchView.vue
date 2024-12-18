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
        const response = await axios.get(`/page?text=${answer.value}`);
        pages.value = response.data.data;
      } catch (err) {
        console.error(err);
      }
    };

    const handleSearch = () => {
      if (answer.value == null || answer.value.length <= 0) {
        return;
      }
      router.replace({query: {answer: answer.value}});
      fetchData();
    }

    const goHome = () => {
      router.push({name: 'home'});
    }

    onMounted(handleSearch);

    return {
      answer,
      pages,
      handleSearch,
      goHome,
    };
  },
};
</script>

<template>
  <el-container>
    <el-header>
      <el-row style="margin-top:50px; margin-left: 20px">
        <img alt="logo"
             src="@/assets/logo.svg"
             style="width: 40px; height: 40px;"
             @click="goHome"/>
        <el-input v-model="answer"
                  clearable
                  placeholder="请输入内容"
                  style="width: 500px; height: 40px; margin-left: 10px; font-size: 20px"
                  @keydown.enter="handleSearch">
        </el-input>
        <el-button style="height: 40px; margin-left: 5px"
                   type="primary"
                   @click="handleSearch">
          <strong>Search</strong>
        </el-button>
      </el-row>
      <el-divider/>
    </el-header>
    <el-main style="margin-top:50px; margin-left: 50px">
      <div v-for="page in pages">
        <el-row>
          <a :href="page.url" style="margin-top: 10px" target="_blank">
            <el-text v-if="page.keywords"
                    v-highlight="[answer, 'highlight-class']"
                    style="font-size: 20px"
                    v-html="page.keywords"
                    truncated>
            </el-text>
            <el-text v-else
                    v-highlight="[answer, 'highlight-class']"
                    style="font-size: 20px"
                    v-html="page.title"
                    truncated>
            </el-text>
          </a>
        </el-row>
        <el-row>
          <el-text v-highlight="[answer, 'highlight-class']"
                   style="margin-top:10px; font-size: 16px"
                   v-html="page.describe"
                   truncated>
          </el-text>
        </el-row>
        <el-row>
          <a :href="page.url"
             style="margin-top:10px; font-size: 12px"
             target="_blank">
            <el-text truncated>{{ page.url }}</el-text>
          </a>
        </el-row>
        <el-divider style="margin: 10px 0"/>
      </div>
    </el-main>
    <el-footer>
    </el-footer>
  </el-container>
</template>
<style scoped>
</style>
