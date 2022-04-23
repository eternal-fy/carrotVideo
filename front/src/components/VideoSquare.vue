<template>
  <div class="videoList-item">
    <router-link :to='{path:"/videoShow/id",query:{id:this.videoResource}}' target="_blank" class="videoitem">
      <div class="videoitem-img"><img class="videoitem-img-img" :src="imgResource"></div>
      <div class="videoitem-botton">{{ infoResource }}</div>
    </router-link>
  </div>

</template>

<script>
export default {
  name: "VideoSquare",
  props: {
    videoResource: String,
    imgResource: String,
    infoResource: String,
  },
  methods:{
    /*routeTo:function (){
      this.$router.push({ name: 'videoShow', query: { id: this.videoResource }})
    }*/
  }
}
</script>

<style scoped>
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
</style>