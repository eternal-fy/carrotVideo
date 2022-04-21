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
              >登陆
              </el-button
              >
            </el-form-item>
            <el-form-item label="thrid-part" >
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
      }
    }
  },
  created() {
  /*  let script = document.createElement('script')
    script.type = 'text/javascript'
    script.src = 'http://connect.qq.com/qc_jssdk.js'
    script.setAttribute("data-appid", "102005140")
    script.setAttribute("data-redirecturi", "http://www.eternalfy.site/auth")
    document.getElementsByTagName('head')[0].appendChild(script)
    script.onload = function () {
      /!* eslint-disable *!/
      QC.Login({
        btnId: "qqLoginBtn"	//插入按钮的节点id
      },function(reqData, opts){//登录成功
        //根据返回数据，更换按钮显示状态方法
        var dom = document.getElementById(opts['btnId']),
            _logoutTemplate=[
              //头像
              '<span><img src="{figureurl}" class="{size_key}"/></span>',
              //昵称
              '<span>{nickname}</span>',
              //退出
              '<span><a href="javascript:QC.Login.signOut();">退出</a></span>'
            ].join("");
        dom && (dom.innerHTML = QC.String.format(_logoutTemplate, {
          nickname : QC.String.escHTML(reqData.nickname), //做xss过滤
          figureurl : reqData.figureurl
        }));
      }, function(opts){//注销成功
        alert('QQ登录 注销成功');
      }
    )
      /!* eslint-enable *!/
    }*/

  },
  methods: {
    userLogin: function () {
      /*      let username = this.userInfo.username
            let password = this.userInfo.password*/
      this.$http({
            url: "/entry/login",
            method: "get",
            data: {
              name: 'ld'
            },
            header: {"Content-Type": "application/json"}
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
        console.log(ref.data)
        window.open(ref.data)
      });

      alert("qq登陆成功")
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
</style>