<template>
  <div class="container">
    <el-row>
      <el-col :span="12">
        <el-button style="width: 100px;">
          <nuxt-link to="/create">
          新規作成
          </nuxt-link>
        </el-button>
      </el-col>
      <el-col :span="12">
        <el-button style="width: 100px;">編集</el-button>
      </el-col>
      <itemTable :tableData="data"/>
    </el-row>
  </div>
</template>

<script>
import itemTable from "~/components/table.vue"

export default {
  components: {
    itemTable
  },
  async asyncData({ app }) {
    let data = await app.$axios.$get("/api")
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
  width: 80%;
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

a {
  text-decoration: none;
  color: #606266;
}
</style>
