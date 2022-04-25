<template>
    <div id="bar">
      <div class="logoArea"><a class="logo" href="/"></a></div>
      <div class="context">
        <div class="columnContxt">
        </div>
        <div class="loginArea">
          <el-link v-show="!islogin" :underline="false" type="info" href="javascript:void(0)" @click="loginFormSwitch">
            <div>点击登陆<img width="50" height="50" src="../assets/personalCenter.svg"/></div>
          </el-link>
          <router-link v-show="islogin" to="/personalCenter" float="right" target="_blank">个人中心2<img width="50"
                                                                                                     height="50"
                                                                                                     src="../assets/personalCenter.svg"/>
          </router-link>
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
  methods: {
    loginFormSwitch: function () {
      this.loginFormShow = !this.loginFormShow
    },
    loginOn: function (flag) {
      this.loginFormSwitch()
      this.islogin = flag
    }
  }
}
</script>

<style scoped>
#bar {
  background-color: #ffffff;
  display: flex;
  flex-direction: row;
  position: fixed;
  width: 100%;
  height: 90px;
  z-index: 888;
  line-height: 50px;
  overflow: hidden;
  border-bottom: 1px solid rgba(0, 0, 0, 0.11);
}

.context {
  display: flex;
  flex-direction: row;
  padding: 20px 50px;
}

.columnContxt {
  width: 800px;

}

.loginArea {
  width: 500px;
  line-height: 50px;
  font-max-size: 16px;
}

.logoArea {
  padding-left: 20px;
  height: 90px;
  min-width: 200px;
}

.logo {
  height: 100%;
  width: 100%;
  background-image: url("../assets/carrot.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  display: block;
}
</style>