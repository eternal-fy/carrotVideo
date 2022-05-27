<template>
  <div class="profile">
    <div class="profileImg">
      <div class="uinfo-head">
        <el-upload
            accept=".jpg,.png,.jpeg,.gif,.JPG,.JPEG"
            class="upload-demo"
            :http-request="imgComplete"
            :limit="1"
        >
          <img :src="profileImgUrl" width="100" height="100"/>
        </el-upload>
      </div>
    </div>
    <div class="changeHead">
      <el-button type="primary">submit</el-button>
    </div>
      <div>
        <h1 class="uinfo-name">{{ userInfo.Name }}</h1>
        <div class="uinfo-info"></div>
        <hr/>
      </div>
<!--      <ul class="uinfo-desc">
        <li><h3>粉丝</h3>
          <h1>1</h1></li>
        <li><h3>关注</h3>
          <h1>1</h1></li>
      </ul>-->
      <hr>
      <div class="menulist">
        <ul>
          <li class="menu-item">
            <router-link :to='{path:"/personalCenter/personalInformation",query:{username:this.username}}'>我的信息
            </router-link>
          </li>
          <li class="menu-item">
            <router-link :to='{path:"/personalCenter/personalVideos",query:{username:this.username}}'>我的视频
            </router-link>
          </li>
        </ul>
      </div>
    </div>
</template>

<script>
export default {
  name: "Profile",
  mounted() {
    let sendData = new FormData();// 上传文件的data参数
    let username = this.username
    sendData.append('username',username);
    this.$http.post("user/getinformation",sendData).then((res) => {
      if (res.data.Code == 9999) {
        this.userInfo = res.data.TransData
      }
    })
    this.$http.post("user/getuserimgurl",sendData)
        .then((res) => {
          if (res.data.Code == 9999) {
            this.profileImgUrl = res.data.TransData
            return
          }
        })
  },
  props: {
    username: String
  },
  data() {
    return {
      profileImgUrl: '',
      userInfo: Object,
      profileNewImgUrl: ''
    }
  },
  methods: {
    imgComplete: function (data) {
      this.profileNewImg = data.file
      let sendData = new FormData();// 上传文件的data参数
      sendData.append('profileImg', this.profileNewImg);
      this.$http.post("user/postprofileimg", sendData).then((res) => {
        alert(res.data.Msg)
      })

    },
  }

}
</script>

<style scoped>
.profile {
  background: rgb(255, 255, 255);
  border-radius: 8px;
  flex: 0 0 210px;
  max-width: 210px;
  min-width: 210px;
  width: 210px;
  padding-top: 20px;
}

.uinfo-head {
  width: 100px;
  height: 100px;
  margin: 0 auto;
  border-radius: 50%;
  overflow: hidden;
}

.uinfo-name {
  display: block;
  text-align: center;
  font-size: 18px;
  margin: 25px 0;
}

.uinfo-info {
  position: relative;
  display: block;
  min-height: 20px;
  padding: 0 20px;
  font-size: 13px;
  text-align: center;
}

hr {
  border: none;
  height: 1px;
  line-height: 1px;
  font-size: 0;
  color: transparent;
  background-color: rgba(0, 0, 0, .05);
  margin: 0 10px;
}

.uinfo-desc {
  padding: 10px 10px;
  overflow: hidden;
}

.uinfo-desc ul {
  list-style: none;
}

.uinfo-desc li {
  width: 50%;
  padding: 15px 0;
  text-align: center;
  display: inline-block;
  border-radius: 5px;
  cursor: pointer;
}

.uinfo-desc h3 {
  color: #666;
  font-size: 13px;
  margin-bottom: 10px;
}

.uinfo-desc h1 {
  color: #000;
  font-weight: 700;
  font-size: 16px;
}

.uinfo-desc li h1:hover {
  color: red;
}

.menulist ul {
  padding: 20px 8px;
  border: none;
  list-style: none;
}

.menu-item {
  padding: 0 !important;
  text-align: center;
  height: 40px;
  line-height: 40px;
}
.changeHead{
  padding: 0px 70px;
}

</style>