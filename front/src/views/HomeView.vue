<script>
import {Search} from '@element-plus/icons-vue'
import Logo from '@/assets/img/logo_big.png'
import { ref } from 'vue'
import { useVuelidate } from '@vuelidate/core'
import { required } from '@vuelidate/validators'

export default {
  name: "SearchView",
  computed: {
    Logo() {
      return Logo
    },
    Search() {
      return Search
    }
  },
  setup() {
    const answer = ref('')

    const rules = {
      answer: { required }
    }

    const { errors, handleSearch } = useVuelidate(rules, { answer })

    return {
      answer,
      errors,
      handleSearch
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
        <el-col :span="4">
          <el-image :src="Logo"/>
        </el-col>
      </el-row>
      <el-row justify="center">
        <div style="margin-top: 80px">
          <el-input id="answer" name="answer" type="text" style="width: 480px; height: 40px;"  v-model="answer"/>
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
