<template>
  <div class="art-video-player">
      <video class="art-video"   loop controls name="media">
        <source :src="videoUrl" type="video/mp4"/>
      </video>
  </div>

</template>

<script>
export default {
  name: "VideoShow",
  data(){
    return {
      videoUrl:String
    }
  },
  mounted() {
    console.log(this.$route.query.id)
    this.videoUrl=this.$route.query.id
  }
}
</script>

<style scoped>
.art-video-player{
  position: absolute;
  top: 0px;
  right: 0px;
  bottom: 0px;
  left: 0px;
  height: 100%;
  width: 100%;
  background-color: rgb(0, 0, 0);

}
.art-video{
  position: absolute;
  top: 0px;
  right: 0px;
  bottom: 0px;
  left: 0px;
  max-height: 100%;
  max-width: 100%;
  margin: auto;
}
video {
  object-fit: contain;
}

</style>