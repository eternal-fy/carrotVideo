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
              <el-input v-model="userInfo.username" @blur="checkUsername"/>
            </el-form-item>
            <el-form-item label="password">
              <el-input type="password" v-model="userInfo.password" @blur="checkPassword"/>
            </el-form-item>
            <el-form-item label="repassword">
              <el-input type="password" v-model="userInfo.repassword" @blur="checkPassword"/>
            </el-form-item>
            <p class="warning">{{ userInfoWarning.warningMsg }}</p>
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
      },
      userInfoWarning: {
        username: false,
        warningMsg: ''

      }
    }
  },
  methods: {
    userRegister: function () {
      if (this.userInfo.password != this.userInfo.repassword) {
        this.userInfoWarning.warningMsg = '两次输入密码不一致！'
        return;
      }
      if (this.userInfo.password.length < 6) {
        this.userInfoWarning.warningMsg = '密码长度应该大于6'
        return
      }
      if (this.userInfoWarning.username.length < 2) {
        this.userInfoWarning.warningMsg = '用户名太短'
        return
      }
      if (this.userInfoWarning.username == false) {
        this.userInfoWarning.warningMsg = '用户名已被注册'
        return
      }
      this.userInfoWarning.warningMsg = ''
      this.$http.post("/entry/register", {
        username: this.userInfo.username,
        password: this.userInfo.password
      }).then(res => {
            if (res.data.Code == "9999") {
              alert("注册成功！")
              this.$router.push({name: 'index'})
              return
            }
            let msg = res.data.Msg
            alert("注册失败！" + msg)
          }
      )
    },
    checkUsername: function () {
      if (this.userInfo.username.length < 2) {
        this.userInfoWarning.warningMsg = '用户名太短'
        this.userInfoWarning.username = false
        return
      }
      this.$http.post("/entry/checkusername", {
        username: this.userInfo.username
      }).then(res => {
            if (res.data.Code == "0000") {
              this.userInfoWarning.username = false
              this.userInfoWarning.warningMsg = '用户名已被注册'
              return
            }
            this.userInfoWarning.username = true
            this.userInfoWarning.warningMsg = ''
          }
      )

    },
    checkPassword: function (e) {
      let value = e.target.value
      if (value.length < 6) {
        this.userInfoWarning.warningMsg = '密码长度应该大于6'
        return
      }
      this.userInfoWarning.warningMsg = ''
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

.warning {
  color: red;
}
</style>