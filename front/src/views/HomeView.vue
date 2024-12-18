<script>
import {Search} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {useRouter} from 'vue-router';

export default {
  name: "HomeView",
  setup() {
    const answer = ref('');
    const router = useRouter();
    const handleSearch = () => {
      if (answer.value == null || answer.value.length <= 0) {
        return;
      }
      router.push({name: 'search', query: {answer: answer.value}});
    };

    const handleLucky = () => {
      alert("暂时未上线，敬请期待！");
    };

    return {
      answer,
      handleSearch,
      Search,
      handleLucky
    };
  }
};
</script>

<template>
  <el-container>
    <el-header></el-header>
    <el-main style="margin-top: 200px">
      <el-row justify="center">
        <img src="@/assets/logo.svg"/>
      </el-row>
      <el-row justify="center">
        <div style="margin-top: 80px">
          <el-input
            id="answer"
            v-model="answer"
            name="answer"
            style="width: 500px; height: 50px; font-size: 20px"
            type="text"
            @keydown.enter="handleSearch"
          />
        </div>
      </el-row>
      <el-row justify="center">
        <div style="margin-top: 30px">
          <el-button
            :icon="Search"
            plain
            size="large"
            type="primary"
            @click="handleSearch">
            <strong>Good Search</strong>
          </el-button>
          <el-button plain
                     size="large"
                     @click="handleLucky">
            <strong>I'm Feeling Lucky</strong>
          </el-button>
        </div>
      </el-row>
    </el-main>
    <el-footer></el-footer>
  </el-container>
</template>

<style scoped>
</style>
