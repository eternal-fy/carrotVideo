<template>
  <div class="videoArea">
    <div>
      <VideoSquare v-for="(item,index) in videoData" :key="index" :videoObject="item"
      ></VideoSquare>
    </div>
    <hr/>
    <div class="upload-area">
      <div>
        <el-upload
            class="upload-demo"
            accept=".mp4,.avi,.wmv,.mpeg,.m4v,.mov,.MP4,.AVI,.WMV,.MPEG,.M4V,.MOV"
            drag
            limit="1"
            :http-request="fileComplete"
        >
          <el-icon class="el-icon--upload">
            <upload-filled/>
          </el-icon>
          <div class="el-upload__text">
            Drop file here or <em>click to upload</em>
          </div>
        </el-upload>
        <el-button @click="uploadFiles">upload</el-button>
      </div>
      <div class="videoInfo">
        <el-form :model="videoInfo" label-width="120px">
          <el-form-item label="title" :rules="[
        {
          required: true,
          message: 'Please input  title',
          trigger: 'blur',
        }]">
            <el-input v-model="videoInfo.title"/>
          </el-form-item>
          <el-form-item label="type" :rules="[
        {
          required: true,
          message: 'Please select',
          trigger: 'blur',
        }]">
            <el-select v-model="videoInfo.type" placeholder="please select your type">
              <el-option v-for="(item,index) in types" :key="index" :label="item" :value="item"/>

            </el-select>
          </el-form-item>
          <el-form-item label="description">
            <el-input v-model="videoInfo.description"/>
          </el-form-item>
          <el-upload
              accept=".jpg,.png,.jpeg,.gif,.JPG,.JPEG"
              class="upload-demo"
              :http-request="imgComplete"
              :limit="1"
          >
            <el-button type="primary">选择封面图片</el-button>
          </el-upload>

        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import {UploadFilled} from '@element-plus/icons-vue'
import VideoSquare from "../VideoSquare";
import AllType from "@/config/catalog.json"

export default {
  name: "PersonalVideos",
  data() {
    return {
      localUrl: 'file/uploadfiles',
      videoData: Object,
      types: [],
      videoInfo: {
        videoFile: null,
        videoImg: null,
        type: '',
        title: '',
        description: ''
      }
    }
  },
  mounted() {
    for (let i = 1; i < AllType.data.length; i++) {
      this.types.push(AllType.data[i].indexName)
    }
    let cookie = this.$getCookie("name")
    this.$http.post("video/getvideoinfos", {
          "videoType": "index",
          "username": cookie,
          "page": 1,
        }
    ).then(res => {
      this.videoData = res.data
    })
  },
  methods: {
    fileComplete: function (data) {
      this.videoInfo.videoFile = data.file
    },
    imgComplete: function (data) {
      this.videoInfo.videoImg = data.file
    },
    uploadFiles: function () {
      if (this.videoInfo.videoFile == null) {
        alert('请上传视频文件')
      }
      let url = this.localUrl
      let sendData = new FormData();// 上传文件的data参数
      sendData.append('videoFile', this.videoInfo.videoFile);
      sendData.append('videoImg', this.videoInfo.videoImg);
      sendData.append('type', this.videoInfo.type);
      sendData.append('title', this.videoInfo.title);
      sendData.append('description', this.videoInfo.description);
      this.$http.post(url, sendData).then((res) => {
        alert(res.data.Msg)
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
  display: flex;
  flex-direction: row;
}

.videoArea {
  padding: 0px 200px;
}

.videoInfo {
  width: 300px;
}
</style>