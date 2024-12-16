<script>
import {useRoute, useRouter} from 'vue-router';
import axios from 'axios';
import {onMounted, ref, watch} from 'vue';

export default {
  name: "SearchView",
  setup() {
    const route = useRoute();
    const router = useRouter();
    const initialAnswer = route.query.answer || ''; // 提供默认值以防没有查询参数
    const answer = ref(initialAnswer); // 使用ref来创建响应式数据
    const pages = ref([]); // 初始化pages为数组
    const loading = ref(true); // 添加加载状态
    const error = ref(null); // 添加错误状态

    const fetchData = async () => {
      try {
        const url = `http://localhost:8080/page?text=${answer.value}`;
        const response = await axios.get(url);
        pages.value = response.data.data; // 假设响应格式是 { data: [...] }
      } catch (err) {
        console.error(err);
        error.value = '加载数据失败'; // 设置错误消息
      } finally {
        loading.value = false; // 无论成功还是失败，都设置加载状态为false
      }
    };

    onMounted(fetchData);

    // 监听answer的变化并重新获取数据（可选）
    watch(answer, (newVal) => {
      router.replace({query: {answer: newVal}}); // 更新URL查询参数
      fetchData(); // 重新获取数据
    });

    return {
      answer,
      pages,
      loading,
      error,
    };
  },
};
</script>

<template>
  <el-container>
    <el-header>
      <el-row style="margin-top:50px; margin-left: 20px">
        <img src="@/assets/logo.svg" style="width: 40px; height: 40px;"/>
        <el-input style="width: 480px; height: 40px; margin-left: 10px" v-model="answer"/>
      </el-row>
      <el-row>
        <el-divider/>
      </el-row>
    </el-header>
    <el-main style="margin-top:50px; margin-left: 20px">
      <div v-for="page in pages">
        <el-row style="margin-top: 15px">
          <el-link :href="page.url" target="_blank">
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
