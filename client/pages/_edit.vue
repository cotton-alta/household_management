<template>
  <div class="container">
    <el-row>
      <el-col :span="24">
        <span>タイトル</span>
        <el-input placeholder="Please input" v-model="data.Body"></el-input>
      </el-col>
      <el-col :span="24">
        <span>金額</span>
        <el-radio v-model="radio" label="1">出金</el-radio>
        <el-radio v-model="radio" label="2"> 入金</el-radio>
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
        data = await app.$axios.$get(`/api/items/${urlParameter}`),
        radio = "1"
    if(data.Amount < 0) {
      radio = "2"
      data.Amount = -data.Amount
      console.log("入金")
    } else {
      console.log("出金")
    }
    return { data, radio }
  },
  data() {
    return { 
      urlParameter: encodeURIComponent(this.$route.params.edit)
    }
  },
  methods: {
    updateData: function() {
      if(this.radio === "2") {
        this.data.Amount = Number(-this.data.Amount)
        console.log("this.data.Amount", this.data.Amount)
      }
      axios.put(`/api/items/${this.urlParameter}`,
        { title: this.data.Body, genre: Number(this.data.Genre), amount: Number(this.data.Amount) },
        { headers: { "Content-Type": "application/json" } }
      )
        .then(result => {
          console.log(result)
        })
    }
  }
}
</script>

<style scoped>

.el-col {
  width: 100%;
  text-align: left;
}

.el-radio {
  margin: 10px 10px;
}

span {
  display: block;
  width: 100%;
  margin-bottom: 10px;
}

</style>