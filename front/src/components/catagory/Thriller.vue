<template>
  <div>
    <VideoSquare v-for="(item,index) in videoData" :key="index" :videoObject="item"
    ></VideoSquare>
  </div>
</template>

<script>
import VideoSquare from "@/components/VideoSquare";

export default {
  name: "Thriller",
  mounted() {
    this.$http.post("video/getvideoinfos", {
          "videoType": "thriller",
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