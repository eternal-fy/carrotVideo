<template>
  <div class="videoArea">
    <div>
      <VideoSquare v-for="(item,index) in videoData" :key="index" :video-resource="item.VideoResource"
                   :img-resource="item.ImgResource" :info-resource="item.InfoResource"
      ></VideoSquare>
    </div>
    <hr/>
    <div class="upload-area">
      <el-upload
          class="upload-demo"
          drag
          :http-request="fileComplete"
      >
        <el-icon class="el-icon--upload">
          <upload-filled/>
        </el-icon>
        <div class="el-upload__text">
          Drop file here or <em>click to upload</em>
        </div>
        <template #tip>
          <div class="el-upload__tip">
            jpg/png files with a size less than 500kb
          </div>
        </template>
      </el-upload>
      <el-button @click="uploadFiles">upload</el-button>
    </div>
  </div>
</template>

<script>
import {UploadFilled} from '@element-plus/icons-vue'
import VideoSquare from "../VideoSquare";

export default {
  name: "PersonalVideos",
  data() {
    return {
      uploadUrl: 'https://eternalfy.site/api/file/uploadfiles',
      localUrl: '/file/uploadfiles',
      file: null,
      videoData: Object
    }
  },
  mounted() {
    this.$http.post("/video/getvideoinfos", {
          "videoType": "index"
        }
    ).then(res => {
      console.log(res.data)
      this.videoData = res.data
    })
  },
  methods: {
    fileComplete: function (data) {
      if (this.file == null) {
        this.file = data.file
      }
    },
    uploadFiles: function () {
      let url = this.localUrl
      let accessUrl = '/' + this.file.name;//设置上传的访问路径
      let sendData = new FormData();// 上传文件的data参数
      sendData.append('file', this.file);
      sendData.append("accessUrl", accessUrl)
      console.log(sendData);
      this.$http.post(url, sendData).then((res) => {
        console.log(res)
      })
    }
  },
  components: {
    UploadFilled, VideoSquare
  }
}
</script>

<style scoped>
.upload-area {

}

.videoArea{
  padding: 0px 200px;
}
</style>