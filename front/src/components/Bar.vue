<template>
  <div>
  <div id="bar">
    <div class="logoArea"><a  class="logo" href="/"></a></div>
    <div class="context">
      <div class="columnContxt">
      </div>
      <div class="loginArea">
            <el-link :underline="false" type="info" href="javascript:void(0)" @click="loginSwitch">个人中心
              <img width="50" height="50" src="../assets/personalCenter.svg"/>
            </el-link>
      </div>
    </div>
  </div>
  <LoginPage :loginSwitch="loginShow" @showSwitch="loginSwitch"></LoginPage>
  </div>
</template>

<script>
import LoginPage from "@/components/entrance/LoginPage";
export default {
  name: "Column",
  data(){
    return{
      loginShow:false
    }
  },
  components:{
    LoginPage
  },
  methods:{
    loginSwitch:function (){
      this.loginShow = !this.loginShow
    }
  }
}
</script>

<style scoped>
#bar {
  display: flex;
  flex-direction: row;
  position: fixed;
  width: 100%;
  height: 90px;
  z-index: 10;
  line-height: 50px;
  overflow: hidden;
  border-bottom: 1px solid rgba(0, 0, 0, 0.11);
}

.context {
  display: flex;
  flex-direction: row;
  padding: 20px 50px;
}
.columnContxt{
  width: 800px;

}

.loginArea {
  width: 500px;
  line-height: 50px;
  font-max-size: 16px;
}

.logoArea{
  padding-left: 20px;
  height: 90px;
  min-width: 200px;
}
.logo{
  height:100%;
  width: 100%;
  background-image: url("../assets/carrot.png");
  background-size:  100% 100%;
  background-repeat: no-repeat;
  display: block;
}
</style>