<script>
import {Search} from '@element-plus/icons-vue'
import Logo from '@/assets/logo.svg'
import {ref} from 'vue'
import {useRouter} from 'vue-router';

export default {
  name: "HomeView",
  setup() {
    const answer = ref('')
    const router = useRouter();
    const handleSearch = function () {
      router.push({name: 'search', query: {answer: answer.value}});
    }
    return {
      answer,
      handleSearch,
      Logo,
      Search
    }
  }
}
</script>

<template>
  <el-container>
    <el-header>
    </el-header>
    <el-main>
      <el-row justify="center">
        <el-image :src="Logo"/>
      </el-row>
      <el-row justify="center">
        <div style="margin-top: 80px">
          <el-input id="answer" name="answer" type="text" style="width: 480px; height: 40px;" v-model="answer"/>
        </div>
      </el-row>
      <el-row justify="center">
        <div style="margin-top: 30px">
          <el-button :icon="Search" plain size="large" @click="handleSearch">Good Search</el-button>
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
