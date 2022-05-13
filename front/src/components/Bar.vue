<template>
  <div id="bar">
    <div class="logoArea"><a class="logo" href="/"></a></div>
    <div class="context">
      <div class="columnContxt">
        <div class="loginArea">
          <el-link v-show="!islogin" :underline="false" type="info" href="javascript:void(0)" @click="loginFormSwitch">
            点击登陆<img height="40" width="30" src="../assets/personalCenter.svg"/>
          </el-link>
          <div v-show="islogin">
            <ul class="menu">
              <p>登陆后的内容</p>
              <li><router-link  to="/personalCenter" target="_blank">个人中心</router-link></li>
              <li><a href="javascript:void(0)" @click="logout">注销</a></li>
            </ul>
          </div>
        </div>
      </div>

    </div>
  </div>
  <LoginPage :loginFormShow="loginFormShow" @showSwitch="loginFormSwitch" @loginOn="loginOn"></LoginPage>
</template>

<script>
import LoginPage from "@/components/entrance/LoginPage";
export default {
  name: "Column",
  data() {
    return {
      islogin: false,
      loginFormShow: false
    }
  },
  components: {
    LoginPage
  },
  mounted(){
    console.log(document.cookie)
   let isLogin = this.$checkCookie("name")
    if (isLogin){
      this.islogin=true
    }
  },
  methods: {
    loginFormSwitch: function () {
      this.loginFormShow = !this.loginFormShow
    },
    loginOn: function (flag) {
      this.loginFormSwitch()
      this.islogin = flag
    },
    logout: function () {
      this.$deleteCookie()
      this.$router.go(0)
    }
  }
}
</script>

<style scoped>
#bar {
  background-color: #ffffff;
  position: fixed;
  width: 100%;
  height: 90px;
  z-index: 888;
  line-height: 50px;
  border-bottom: 1px solid rgba(0, 0, 0, 0.11);
}

.context {
  padding: 20px 50px;
  float: right;
}

.columnContxt {
  height: 100%;
}

.loginArea {
  padding-left: 20px;
  float: right;
  width: 200px;
  line-height: 30px;
  height: 100%;
  font-max-size: 16px;
}

.logoArea {
  padding-left: 20px;
  height: 100%;
  width: 200px;
  float: left;
}

.logo {
  height: 100%;
  width: 100%;
  background-image: url("../assets/carrot.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  display: block;
}

.menu{
  list-style: none;
  line-height: 0px;
  transition: line-height 0.3s;
}
.menu:hover{
  line-height: 20px;
}

a{
  text-decoration: none;
}
</style>