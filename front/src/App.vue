<template>
  <el-container>
    <el-header>
      <div class="header-content">
        <el-logo class="logo" />
        <div class="search-box">
          <el-input v-model="searchQuery" placeholder="请输入关键词" class="input-search" />
          <el-button type="primary" @click="handleSearch">搜索</el-button>
        </div>
      </div>
    </el-header>
    <el-main>
      <!-- 搜索结果展示区域 -->
      <div v-if="searchResults.length > 0" class="search-results">
        <h2>搜索结果</h2>
        <ul>
          <li v-for="result in searchResults" :key="result.id">{{ result.title }}</li>
        </ul>
      </div>
      <div v-else class="no-results">
        <h2>没有找到相关结果</h2>
      </div>
    </el-main>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      searchQuery: '',
      searchResults: []
    };
  },
  methods: {
    handleSearch() {
      // 模拟搜索逻辑，这里可以使用真实的搜索 API
      if (this.searchQuery) {
        // 模拟搜索结果
        this.searchResults = [
          { id: 1, title: '搜索结果 1：' + this.searchQuery },
          { id: 2, title: '搜索结果 2：' + this.searchQuery },
          { id: 3, title: '搜索结果 3：' + this.searchQuery }
        ];
      } else {
        this.searchResults = [];
      }
    }
  }
};
</script>

<style scoped>
.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  height: 60px;
  background-color: #fff;
  border-bottom: 1px solid #dcdcdc;
}

.logo {
  width: 120px;
  height: 40px;
  background: url('https://www.baidu.com/img/baidu_jgylogo3.gif') no-repeat center center;
  background-size: contain;
}

.search-box {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
}

.input-search {
  width: 70%;
  height: 36px;
  border: 1px solid #dcdcdc;
  border-radius: 4px 0 0 4px;
  padding: 0 10px;
  box-sizing: border-box;
}

.el-button {
  height: 36px;
  border-radius: 0 4px 4px 0;
  margin-left: -1px; /* To remove the border gap between input and button */
}

.search-results, .no-results {
  padding: 20px;
}

.search-results ul {
  list-style: none;
  padding: 0;
}

.search-results li {
  margin: 10px 0;
}
</style>