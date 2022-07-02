<template>
  <div class="videoList-item">
    <router-link :to='{path:"/videoShow/id",query:{videoInfo:JSON.stringify(this.videoObject)}}' target="_blank"
                 class="videoitem">
      <div class="videoitem-img">
        <img class="videoitem-img-img" :src="videoObject.ImgResource">
      </div>
      <div class="uinfo-head">
        <img :src="videoObject.AuthorImg" width="50" height="50"/>
      </div>
      <div class="videoitem-botton">
        标题：<span>{{ videoObject.VideoModel.Title }}</span>
      </div>
      <div class="videoitem-botton">
        描述：<span>{{ videoObject.VideoModel.Description }}</span>
      </div>
    </router-link>
  </div>

</template>

<script>
export default {
  name: "VideoSquare",
  props: {
    videoObject: Object
  },
  methods: {}
}
</script>

<style scoped>
.uinfo-head{
  width: 50px;
  height: 50px;
  border-radius: 50%;
  overflow: hidden;
}
.videoList-item {
  display: inline-block;
  width: 25%;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  padding-right: 16px;
  margin-top: 44px;
  vertical-align: top;
}

.videoitem-img {
  padding-bottom: 56.08%;
  height: 0;
  position: relative;
  font-size: 0;
  width: 100%;
  background-size: 100% auto;
  border-radius: 3px;
  overflow: hidden;
  background-repeat: no-repeat;
}

.videoitem-img:hover {
  border-radius: 20px;
}

.videoitem-img-img {
  width: 100%;
}
a{
  text-decoration: none;
  color:black;
}
</style>