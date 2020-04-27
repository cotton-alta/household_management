<template>
  <div class="table-wrapper">
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column prop="Body" label="タイトル"></el-table-column>
      <!-- <el-table-column prop="Id" label="ID"></el-table-column> -->
      <el-table-column prop="Genre" label="ジャンル"></el-table-column>
      <el-table-column prop="Created" label="記入日"></el-table-column>
      <!-- <el-table-column prop="Updated" label="更新日"></el-table-column> -->
      <el-table-column prop="Amount" label="入出金"></el-table-column>
      <el-table-column prop="Balance" label="残高"></el-table-column>
      <el-table-column
        fixed="right"
        label=""
        width="88">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="deleteData(scope.$index, tableData, scope)">削除</el-button>
          <el-button type="text" size="small" @click="updateLink(scope.$index, tableData, scope)">編集</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
  <!-- /.container -->
</template>

<script>
import axios from "axios"

export default {
  props: [ "tableData" ],
  methods: {
    deleteData: function(index, tableData, scope) {
      let rowId = tableData.length - index - 1
      axios.delete(`/api/items/${rowId}`)
        .then(result => {
          this.$router.go({path: this.$router.currentRoute.path, force: true})
        }).catch(err => {
          console.log(err)
        })
    },
    updateLink: function(index, tableData, scope) {
      let rowId = tableData.length - index - 1
      this.$router.push(`/${rowId}`)
    }
  }
}
</script>

<style scoped>

</style>