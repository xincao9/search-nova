<script>
import {useRoute, useRouter} from 'vue-router';
import {ref} from 'vue';

export default {
  name: "SearchView",
  setup() {
    const route = useRoute();
    const router = useRouter();
    const initialAnswer = route.query.answer || '';
    const answer = ref(initialAnswer);
    const pages = ref([
      {
        "title": "凤凰,凤凰网,凤凰新媒体,凤凰卫视,凤凰卫视中文台,资讯台,电影台,欧洲台,美洲台,凤凰周刊",
        "describe": "凤凰网是中国领先的综合门户网站，提供含文图音视频的全方位综合新闻资讯、深度访谈、观点评论、财经产品、互动应用、分享社区等服务，同时与凤凰无线、凤凰宽频形成三屏联动，为全球主流华人提供互联网、无线通信、电视网三网融合无缝衔接的新媒体优质体验。",
        "url": "https://www.ifeng.com"
      },
      {
        "title": "生活故事,架构,大数据,一线,FastDFS,开发者,编程,代码,开源,IT网站",
        "describe": "本站是纯洁的微笑的技术分享博客。内容涵盖生活故事、Java后端技术、Spring Boot、Spring Cloud、微服务架构、大数据演进、高可用架构、中间件使用、系统监控等相关的研究与知识分享。",
        "url": "https://www.ityouknow.com"
      }
    ]);
    const loading = ref(true);
    const error = ref(null);

    // const fetchData = async () => {
    //   try {
    //     const url = `http://localhost:8080/page?text=${answer.value}`;
    //     const response = await axios.get(url);
    //     pages.value = response.data.data;
    //   } catch (err) {
    //     console.error(err);
    //     error.value = '加载数据失败';
    //   } finally {
    //     loading.value = false;
    //   }
    // };
    //
    // onMounted(fetchData);
    //
    // watch(answer, (newVal) => {
    //   router.replace({query: {answer: newVal}});
    //   fetchData();
    // });

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
        <img src="@/assets/logo.svg" style="width: 40px; height: 40px;" alt="logo"/>
        <el-input style="width: 480px; height: 40px; margin-left: 10px" v-model="answer"/>
      </el-row>
      <el-row>
        <el-divider/>
      </el-row>
    </el-header>
    <el-main style="margin-top:50px; margin-left: 50px">
      <div v-for="page in pages">
        <el-row>
          <a style="margin-top: 15px" :href="page.url" target="_blank">
            <text style="font-size: 20px">{{ page.title }}</text>
          </a>
        </el-row>
        <el-row>
          <text style="margin-top:10px; font-size: 16px">{{ page.describe }}</text>
        </el-row>
      </div>
    </el-main>
    <el-footer>
    </el-footer>
  </el-container>
</template>
<style scoped>
</style>
