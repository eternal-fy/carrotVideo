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
              <el-input type="password" v-model="userInfo.password"/>
            </el-form-item>
            <p class="warning">{{ userInfoWarning.warningMsg }}</p>
            <el-form-item>
              <el-button type="primary" @click="userLogin"
              >登陆
              </el-button
              >
            </el-form-item>
            <el-form-item label="thrid-part">
              <a href="javascript:void(0)"><img src="@/assets/qq.png" @click="qqLogin" height="32"/></a>
              <div class="registerArea">
                <router-link to="/register" float="right" target="_blank">立即注册</router-link>
              </div>
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
      },
      userInfoWarning: {
        warningMsg: ''
      }
    }
  },
  methods: {
    userLogin: function () {
      let username = this.userInfo.username
      let password = this.userInfo.password
      if (password.length < 6) {
        this.userInfoWarning.warningMsg = '密码长度应该大于6'
        return
      }
      if (username.length < 2) {
        this.userInfoWarning.warningMsg = '用户名太短'
        return
      }
      this.$http.post("/entry/login", {
            "username": username,
            "password": password
          }
      ).then(ref => {
        console.log(JSON.stringify(ref.data))
      });
      this.$emit("loginComplete", true)
    },
    qqLogin: async function () {
      await this.$http({
            url: "/entry/auth",
            method: "get",
          }
      ).then(ref => {
        window.open(ref.data)
      });
      this.$emit("loginComplete", true)
    }
  },
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

a {
  text-decoration: none;
  color: #ff2626;
  font-size: 14px;
}

.registerArea {
  position: absolute;
  right: 26px;
}

.el-button {
  width: 500px;
}
.warning {
  color: red;
}
</style>