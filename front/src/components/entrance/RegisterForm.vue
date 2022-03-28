<template>
  <div class="registerSquare">
    <div class="registerForm">
      <div class="registerContext">
        <span class="formTitle">注册</span>
        <div class="formContext">
          <el-form
              label-width="100px"
              :model="formLabelAlign"
              style="max-width: 460px"
          >
            <el-form-item label="username">
              <el-input v-model="userInfo.username"/>
            </el-form-item>
            <el-form-item label="password">
              <el-input v-model="userInfo.password"/>
            </el-form-item>
            <el-form-item label="repassword">
              <el-input v-model="userInfo.repassword"/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="userRegister"
              >submit
              </el-button
              >
            </el-form-item>
          </el-form>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: "RegisterForm",
  data() {
    return {
      userInfo: {
        username: '',
        password: '',
        repassword: '',
      }
    }
  },
  methods:{
    userRegister:function (){
      alert(this.userInfo)
    }
  }
}
</script>

<style scoped>
.registerSquare {
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
  position: absolute;
  background-color: white;
  width: 425px;
  border-radius: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
}

.registerForm {
  line-height: 1.5;
  padding: 20px 40px;
  position: relative;
}

.registerContext {
  text-align: center;
}

.formTitle {
  font-size: 20px;
  color: #222;
  font-weight: 600;
}

.formContext {
  display: flex;
  flex-direction: column;
  margin-top: 10px;
}
</style>