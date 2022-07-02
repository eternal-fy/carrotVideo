<template>
  <PersonalBar :title="personalBarTitle"></PersonalBar>
  <div class="art-video-player">
    <div class="art-video-area">
      <video class="art-video" loop controls name="media">
        <source :src="videoInfo.VideoResource" type="video/mp4"/>
      </video>
      <div class="video-operate">
        <div class="starArea" @click="star">
          {{ count }}
          <span>点赞</span>
          <svg width="24" height="24" xmlns="http://www.w3.org/2000/svg">
            <path
                d="M11.215 2.35c2.148-.643 4.132 1.232 3.431 4.662l-.267 1.255h4.084c1.814 0 3.05 1.46 2.778 3.214l-.032.176-1.427 6.748c-.363 1.72-2.045 3.08-3.805 3.08H5.302A2.048 2.048 0 0 1 3.25 19.44v-8.497c0-1.13.914-2.044 2.044-2.046l1.3-.003c.611-.043 1.023-.346 1.479-.965.493-.669.702-1.049.996-1.819l.156-.42.14-.398c.035-.107.07-.222.105-.349l.195-.725c.308-1.075.671-1.605 1.55-1.869zM6.86 10.37c-.16.019-.326.028-.502.027l-1.061.002a.546.546 0 0 0-.548.546v8.497c0 .298.246.544.552.544l1.519-.001zm6.315-3.658c.477-2.332-.482-3.238-1.53-2.924-.274.082-.39.274-.59 1.025l-.203.749a7.854 7.854 0 0 1-.129.393l-.124.353-.133.358A9.184 9.184 0 0 1 9.28 8.82a4.54 4.54 0 0 1-.918.96l-.041 10.204h7.656c1 0 2.014-.78 2.299-1.737l.039-.153 1.426-6.747c.19-.9-.358-1.58-1.278-1.58h-5.02a.75.75 0 0 1-.73-.924l.156-.69z"
                :fill="iconColor"/>
          </svg>
        </div>
        <div class="infoArea">
          <div class="uinfo-head">
            <router-link to="/personalCenter/personalInformation" target="_blank">
              <img :src="userinfo.profileImgUrl" width="50" height="50"/>
            </router-link>
          </div>
          <span>作者：{{ userinfo.name }}</span>
          <span>标题：{{ videoModelInfo.Title }}</span>
          <span>描述：{{ videoModelInfo.Description }}</span>
        </div>
      </div>
      <el-form>
        <el-form-item label="留言" prop="desc">
          <el-input v-model="desc" type="textarea"/>
        </el-form-item>
        <el-button @click="msgSubmit">提交</el-button>
      </el-form>
      <div>
        <msg-box v-for="(item,index) in msgList" :key="index" :msgcontext="item.Desc"
                 :username="item.NickName"></msg-box>
      </div>
    </div>
  </div>
</template>

<script>
import PersonalBar from "./PersonalBar";
import MsgBox from "../MsgBox";

export default {
  name: "VideoShow",
  data() {
    return {
      iconColor: 'black',
      videoInfo: Object,
      videoModelInfo: Object,
      personalBarTitle: '播放中心',
      desc: '',
      msgList: [],
      userinfo: {
        name: '',
        profileImgUrl: '',
      },
      count: 0

    }

  },
  mounted() {
    this.videoInfo = JSON.parse(this.$route.query.videoInfo)
    this.videoModelInfo = this.videoInfo.VideoModel
    let sendData = new FormData();// 上传文件的data参数
    sendData.append('username', this.videoModelInfo.Username);
    this.$http.post("user/getuserimgurl", sendData)
        .then((res) => {
          if (res.data.Code == 9999) {
            this.userinfo.profileImgUrl = res.data.TransData
            return
          }
        })
    this.$http.post("user/getusernickname",
        sendData
    )
        .then((res) => {
          if (res.data.Code == 9999) {
            this.userinfo.name = res.data.TransData
            return
          }
        })
    let msgData = new FormData();// 上传文件的data参数
    msgData.append('videoid', this.videoModelInfo.ID);
    this.$http.post("video/getmsg",
        msgData
    ).then((res) => {
      if (res.data.Code == 9999) {
        this.msgList = res.data.TransData
        console.log(this.msgList)
        return
      }
    })
    if (this.$getCookie("name")!=""){
      let starData = new FormData();
      starData.append("videoid", this.videoModelInfo.ID)
      this.$http.post("video/getstar", starData)
          .then((res) => {
            if (res.data.Code == 9999) {
              let flag = res.data.TransData
              if (flag > 0) {
                this.iconColor = 'red'
              } else {
                this.iconColor = 'black'
              }
              return
            }
          })
    }
    let starCountData = new FormData();
    starCountData.append("videoid", this.videoModelInfo.ID)
    this.$http.post("video/getstarcount", starCountData)
        .then((res) => {
          if (res.data.Code == 9999) {
            this.count = res.data.TransData
            return
          }
        })
  },
  methods: {
    checkLogin:function (){
      let cookie = this.$getCookie("name")
      if(cookie==""){
        alert("please login")
        return false
      }
      return true
    },
    star: function () {
     let flag = this.checkLogin()
      if (!flag){
        return
      }
      this.iconColor = this.iconColor == 'red' ? 'black' : 'red'

      let starData = new FormData();
      starData.append("videoid", this.videoModelInfo.ID)
      starData.append("author", this.videoModelInfo.Username)
      this.$http.post("video/star", starData)
          .then((res) => {
            if (res.data.Code == 9999) {
              let starCountData = new FormData();
              starCountData.append("videoid", this.videoModelInfo.ID)
              this.$http.post("video/getstarcount", starCountData)
                  .then((res) => {
                    if (res.data.Code == 9999) {
                      this.count = res.data.TransData
                      return
                    }
                  })

            }
          })
    },
    msgSubmit: function () {
      let flag = this.checkLogin()
      if (!flag){
        return
      }
      let sendData = new FormData();// 上传文件的data参数
      sendData.append('videoid', this.videoModelInfo.ID);
      sendData.append('desc', this.desc);
      this.$http.post("video/postmsg",
          sendData
      ).then((res) => {
        console.log(res)
        let msgData = new FormData();// 上传文件的data参数
        msgData.append('videoid', this.videoModelInfo.ID);
        this.$http.post("video/getmsg",
            msgData
        ).then((res) => {
          if (res.data.Code == 9999) {
            this.msgList = res.data.TransData
            this.desc = ''
            return
          }
        })
      })

    }
  },
  components: {
    MsgBox,
    PersonalBar
  }
}
</script>

<style scoped>
.art-video-player {
  height: 100%;
  min-width: 1400px;
  padding: 100px 100px;
  background-color: #fafafa;

}

.art-video-area {
  height: 70%;
  width: 70%;
}

.art-video {
  max-height: 100%;
  max-width: 100%;
}

.video-operate {
  background-color: white;
  width: 100%;
}

video {
  object-fit: contain;
}

.art-video-area {

}

.starArea {
  cursor: pointer;
  margin-left: 50%;
  line-height: 30px;
  display: flex;
}

.infoArea {
  display: flex;
  flex-direction: row;
  line-height: 50px;
  background-color: whitesmoke;
}

.uinfo-head {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  overflow: hidden;
}

.infoArea span {
  margin: 0 50px;
}
</style>