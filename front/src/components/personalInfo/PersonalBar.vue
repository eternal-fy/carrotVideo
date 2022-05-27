<template>
  <div id="personalbar">
    <div class="p-logoArea"><a class="p-logo" href="/"></a></div>
    <div class="context">
      <div class="columnContxt">
        <h1>{{title}}</h1>
        <div class="uinfo-head">
          <router-link  to="/personalCenter/personalInformation" target="_blank">
            <img :src="profileImgUrl" width="50" height="50"/>
          </router-link>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
export default {
  name: "PersonalBar",
  props: {
    title: String
  },
  mounted() {
    let sendData = new FormData();// 上传文件的data参数
    let username = this.$getCookie("name")
    sendData.append('username',username);
    this.$http.post("user/getuserimgurl",sendData)
        .then((res) => {
          if (res.data.Code == 9999) {
            this.profileImgUrl = res.data.TransData
            return
          }
        })
  },
  data() {
    return {
      profileImgUrl: ''
    }
  }
}
</script>

<style scoped>
#personalbar {
  background-color: #ffffff;
  position: fixed;
  width: 100%;
  height: 90px;
  z-index: 888;
  line-height: 50px;
  border-bottom: 1px solid rgba(0, 0, 0, 0.11);
}

.context {
  float: right;
  height: 100%;
}

.columnContxt {
  height: 100%;
  display: flex;
  flex-direction: row-reverse;
}

.p-logoArea {
  padding-left: 20px;
  height: 100%;
  width: 200px;
  float: left;
}

.p-logo {
  height: 100%;
  width: 100%;
  background-image: url("../../assets/carrot.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  display: block;
}

a {
  text-decoration: none;
}

.uinfo-head {
  height: 50px;
  border-radius: 50%;
  overflow: hidden;
}

img {
  cursor: pointer;
}
</style>