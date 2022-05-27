<template>
  <div id="bar">
    <div class="logoArea"><a class="logo" href="/"></a></div>
    <div class="context">
      <div class="columnContxt">
        <div class="loginArea">
          <el-link v-show="!islogin" :underline="false" type="info" href="javascript:void(0)" @click="loginFormSwitch">
            点击登录<img height="40" width="30" src="../assets/personalCenter.svg"/>
          </el-link>
          <div v-show="islogin" class="face">
            <div class="uinfo-head">
              <router-link class="menu-item" :to='{path:"/personalCenter",query:{username:this.userInfo.username}}' target="_blank">
                <img :src="userInfo.profileImgUrl" width="50" height="50"/>
              </router-link>
            </div>
            <div class="logout">
              <el-button @click="logout">注销</el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <LoginPage :loginFormShow="loginFormShow" @showSwitch="loginFormSwitch"></LoginPage>
</template>

<script>
import LoginPage from "@/components/entrance/LoginPage";

export default {
  name: "Column",
  data() {
    return {
      islogin: false,
      loginFormShow: false,
      userInfo: {
        username:'',
        name: '',
        age: Number,
        gender: Number,
        level: Number,
        address: '',
        phone: '',
        profileImgUrl: ''
      }
    }
  },
  components: {
    LoginPage
  },
  mounted() {
    let sendData = new FormData();// 上传文件的data参数
    let username = this.$getCookie("name")
    sendData.append('username',username);
    this.$http.post("user/getinformation",sendData)
        .then((res) => {
          if (res.data.Code == 9999) {
            this.islogin = true
            let data = res.data.TransData
            this.userInfo.username=data.Username
            this.userInfo.name = data.Name
            this.userInfo.age = data.Age
            this.userInfo.gender = data.Gender
            this.userInfo.level = data.Level
            this.userInfo.address = data.Address
            this.userInfo.phone = data.Phone
          }
        })
    this.$http.post("user/getuserimgurl",sendData)
        .then((res) => {
          if (res.data.Code == 9999) {
            this.islogin = true
            this.userInfo.profileImgUrl = res.data.TransData
            return
          }
          this.islogin = false
        })


  },
  methods: {
    loginFormSwitch: function () {
      this.loginFormShow = !this.loginFormShow
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
  display: flex;
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

a {
  text-decoration: none;
}


.uinfo-head {
  width: 50px;
  height: 50px;
  margin: 0 auto;
  border-radius: 50%;
  overflow: hidden;
}

img {
  cursor: pointer;
}

.face {
  display: flex;
  flex-direction: row;
}

.logout {
  line-height: 50px;
}
</style>