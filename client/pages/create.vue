<template>
  <div class="container">
    <el-row>
      <el-col :span="24">
        <span>タイトル</span>
        <el-input placeholder="Please input" v-model="title"></el-input>
      </el-col>
      <el-col :span="24">
        <span>金額</span>
        <el-input placeholder="Please input" v-model="amount" ></el-input>
      </el-col>
      <el-col :span="24">
        <span>ジャンル</span>
        <el-select style="width: 100%;" v-model="genre" placeholder="please select your zone">
          <el-option label="ジャンル1" value="1"></el-option>
          <el-option label="ジャンル2" value="2"></el-option>
        </el-select>
      </el-col>
    </el-row>
    <el-row>
      <el-button style="width: 100px;" @click="postData">作成</el-button>
      <el-button style="width: 100px;">
        <nuxt-link to="/">戻る</nuxt-link>
      </el-button>
    </el-row>
  </div>  
</template>

<script>
import axios from "axios"

export default {
  data() {
    return { 
      title: "",
      genre: "",
      amount: "",

    }
  },
  methods: {
    postData: function() {
      axios.post("/api/items",
        { title: this.title, genre: Number(this.genre), amount: Number(this.amount) },
        { headers: { "Content-Type": "application/json" } }
      )
    }
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

.el-col {
  width: 100%;
  text-align: left;
}
</style>