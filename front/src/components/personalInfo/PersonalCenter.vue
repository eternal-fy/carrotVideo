<template>
    <PersonalBar></PersonalBar>
    <div class="content-main">
      <Profile></Profile>
      <router-view></router-view>
      <PersonalInformation></PersonalInformation>
    </div>
    <div>
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
</template>


<script>
import {UploadFilled} from '@element-plus/icons-vue'
import PersonalBar from "./PersonalBar";
import Profile from "./Profile";
import PersonalInformation from "./PersonalInformation";

export default {
  name: "PersonalCenter",
  data() {
    return {
      uploadUrl: 'https://eternalfy.site/api/file/uploadfiles',
      localUrl: '/file/uploadfiles',
      file: null
    }
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
    PersonalInformation,
    Profile,
    PersonalBar,
    UploadFilled
  },
}
</script>
<style scoped>
.content-main {
  display: flex;
  width: 100%;
  height: 500px;
  margin:0  auto;
  padding: 130px 50px;
  background-color: #f7f7f7;
}
</style>

