<template>
  <div>
    <el-form>
      <el-form-item :label="username" prop="desc" label-width="120px">
        <el-input  type="textarea" :value="msgcontext"/>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "MsgBox",
  props:{
    username:String,
    msgcontext:String
  }
}
</script>

<style scoped>

</style>