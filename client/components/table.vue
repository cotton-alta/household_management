<template>
  <div class="table-wrapper">
    <el-table
      :data="tableData"
      border
      style="width: 100%">
      <el-table-column prop="Id" label="ID"></el-table-column>
      <el-table-column prop="Created" label="記入日"></el-table-column>
      <el-table-column prop="Updated" label="更新日"></el-table-column>
      <el-table-column prop="Amount" label="金額"></el-table-column>
      <el-table-column prop="Body" label="備考"></el-table-column>
      <el-table-column prop="Balance" label="残高"></el-table-column>
      <el-table-column prop="Genre" label="ジャンル"></el-table-column>
    </el-table>
  </div>
  <!-- /.container -->
</template>

<script>
import axios from "axios"

export default {
  props: [ "tableData" ]
}
</script>

<style scoped>

</style>