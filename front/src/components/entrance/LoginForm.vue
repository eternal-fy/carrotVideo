<template>
  <div class="loginSquare">
    <div class="loginForm">
      <div class="loginContext">
        <span class="formTitle">登录</span>
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
            <el-form-item>
              <el-button type="primary" @click="userLogin"
              >submit
              </el-button
              >
            </el-form-item>
            <el-form-item label="thrid-part">
              <a href="javascript:void(0)" @click="qqLogin"><img src="@/assets/qq.png" height="32"/></a>

            </el-form-item>
          </el-form>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: "LoginForm",
  data() {
    return {
      userInfo: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    userLogin: function () {
      let username = this.userInfo.username
      let password = this.userInfo.password
      alert(username + "-" + password + "登陆成功")
    },
    qqLogin: function () {
      alert("qq登陆成功")
    }
  }
}
</script>

<style scoped>
.loginSquare {
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
  position: absolute;
  background-color: white;
  width: 425px;
  border-radius: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
}

.loginForm {
  line-height: 1.5;
  padding: 20px 40px;
  position: relative;
}

.loginContext {
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