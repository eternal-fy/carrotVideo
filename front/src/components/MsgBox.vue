<template>
  <div>
    <el-form>
      <p>{{createtime}}</p>
      <el-form-item :label="username" prop="desc" label-width="120px">
        <el-input  type="textarea" :value="msgcontext" readonly />
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "MsgBox",
  props:{
    username:String,
    msgcontext:String,
    createtime:String
  },
}
</script>

<style scoped>
input[type="textarea"]{
  resize:none;
}
</style>