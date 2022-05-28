<template>
  <div>
    <VideoSquare v-for="(item,index) in videoData" :key="index" :videoObject="item"
    ></VideoSquare>
  </div>
</template>

<script>
import VideoSquare from "@/components/VideoSquare";

export default {
  name: "Adventure",
  mounted() {
    this.$http.post("video/getvideoinfos", {
          "videoType": "adventure",
          "page": 1,
        }
    ).then(res => {
      this.videoData = res.data
    })
  },
  data() {
    return {
      videoData: Object,
    }
  },
  components: {
    VideoSquare
  }
}
</script>

<style scoped>

</style>