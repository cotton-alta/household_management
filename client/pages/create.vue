<template>
  <div class="container">
    <el-row>
      <el-col :span="24">
        <span>内容</span>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </el-col>
      <el-col :span="24">
        <span>金額</span>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </el-col>
      <el-col :span="24">
        <span>ジャンル</span>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </el-col>
    </el-row>
    <el-row>
      <el-button style="width: 100px;">作成</el-button>
      <el-button style="width: 100px;">
        <nuxt-link to="/">戻る</nuxt-link>
      </el-button>
    </el-row>
  </div>  
</template>

<script>
export default {
  data() {
    return { input: "" }
  }
}
</script>

<style scoped>
.container {
  display: block;
  width: 90%;
  max-width: 800px;
  margin: 0 auto;
  min-height: 100vh;
}
</style>