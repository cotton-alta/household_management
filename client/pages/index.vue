<template>
  <div class="container">
    <el-row>
      <el-col :span="24">
        <span>タイトルから検索</span>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </el-col>
      <el-col>
      <el-collapse>
        <el-collapse-item title="絞り込み" name="1">
          <span>ジャンル</span>
          <el-input placeholder="Please input" v-model="input"></el-input>
        </el-collapse-item>
      </el-collapse>
      </el-col>
      <el-col :span="24">
        <el-button style="width: 200px;">
          <nuxt-link to="/create">
          新規作成
          </nuxt-link>
        </el-button>
      </el-col>
      <itemTable :tableData="data"/>
    </el-row>
  </div>
</template>

<script>
import itemTable from "~/components/table.vue"
import moment from "moment"

export default {
  components: {
    itemTable
  },
  async asyncData({ app }) {
    let data = await app.$axios.$get("/api")
    data.map(item => {
      if(item.Amount != 0) {
        item.Amount = -item.Amount
        if(item.Amount > 0) {
          item.Amount = "+" + item.Amount
        }
      }
      item.Created = moment(item.Created).format('L')
      item.Updated = moment(item.Updated).format('L')
    })
    return { data }
  },
  data() {
    return {
      input: ''
    }
  }
}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  width: 95%;
  max-width: 1000px;
}

.el-row {
  margin: 100px auto;
  text-align: center;
  width: 100%;
}

.el-col {
  border-radius: 4px;
  margin: 30px 0px;
}

.el-button {
  position: relative;
  height: 40px;
}

a {
  line-height: 40px;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  text-decoration: none;
  color: #606266;
}
</style>
