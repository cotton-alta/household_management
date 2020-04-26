<template>
  <div class="container">
    <el-row>
      <el-col :span="24">
        <span>タイトル</span>
        <el-input placeholder="Please input" v-model="data.Body"></el-input>
      </el-col>
      <el-col :span="24">
        <span>金額</span>
        <el-input placeholder="Please input" v-model="data.Amount" ></el-input>
      </el-col>
      <el-col :span="24">
        <span>ジャンル</span>
        <el-select style="width: 100%;" v-model="data.Genre" placeholder="please select your zone">
          <el-option label="ジャンル1" value="1"></el-option>
          <el-option label="ジャンル2" value="2"></el-option>
        </el-select>
      </el-col>
    </el-row>
    <el-row>
      <el-button style="width: 100px;" @click="deleteData">削除</el-button>
      <el-button style="width: 100px;" @click="updateData">編集完了</el-button>
      <el-button style="width: 100px;">
        <nuxt-link to="/">戻る</nuxt-link>
      </el-button>
    </el-row>
  </div>
  <!-- /.container -->
</template>

<script>
import axios from "axios"

export default {
  async asyncData({ app, route }) {
    let urlParameter = encodeURIComponent(route.params.edit),
        data = await app.$axios.$get(`/api/items/${urlParameter}`)
    return { data }
  },
  data() {
    return { urlParameter: encodeURIComponent(this.$route.params.edit) }
  },
  methods: {
    updateData: function() {
      axios.post(`/api/items/${this.urlParameter}`,
        { title: this.data.Body, genre: Number(this.data.Genre), amount: Number(this.data.Amount) },
        { headers: { "Content-Type": "application/json" } }
      )
    },
    deleteData: function() {
      console.log(this.urlParameter)
      axios.delete(`/api/items/${this.urlParameter}`)
        .then(result => {
          this.$router.push("/")
        }).catch(err => {
          console.log(err)
        })
    }
  }
}
</script>

<style scoped>

</style>