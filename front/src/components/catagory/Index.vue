<template>
  <div>
    <VideoSquare v-for="(item,index) in videoData" :key="index" :video-resource= "item.VideoResource" :img-resource="item.ImgResource" :info-resource="item.InfoResource"
                ></VideoSquare>
  </div>
</template>

<script>
import VideoSquare from "@/components/VideoSquare";

export default {
  name: "Index",
  mounted(){
    this.$http.post( "/video/getvideoinfos",{
        "videoType": "index"
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