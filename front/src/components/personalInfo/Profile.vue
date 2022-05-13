<template>
<div class="profile">
  <div class="profileImg">
    <div class="uinfo-head">
      <img src="../../assets/face.jpg" width="100" height="100"/>
    </div>
    <h1 class="uinfo-name">ice、wind</h1>
    <div class="uinfo-info"></div>
    <hr>
  </div>
  <ul class="uinfo-desc">
    <li><h3>粉丝</h3><h1>1</h1></li>
    <li><h3>关注</h3><h1>1</h1></li>
  </ul>
  <hr>
  <div class="menulist">
    <ul>
      <li class="menu-item">我的信息</li>
      <li class="menu-item">我的视频</li>
    </ul>
  </div>

</div>
</template>

<script>
export default {
  name: "Profile"
}
</script>

<style scoped>
.profile{
  background: rgb(255, 255, 255);
  border-radius: 8px;
  flex: 0 0 210px;
  max-width: 210px;
  min-width: 210px;
  width: 210px;
  padding-top: 20px;
}

.uinfo-head{
  width: 100px;
  height: 100px;
  margin: 0 auto;
  border-radius: 50%;
  overflow: hidden;
}
.uinfo-name{
  display: block;
  text-align: center;
  font-size: 18px;
  margin: 25px 0;
}
.uinfo-info{
  position: relative;
  display: block;
  min-height: 20px;
  padding: 0 20px;
  font-size: 13px;
  text-align: center;
}
hr{
  border: none;
  height: 1px;
  line-height: 1px;
  font-size: 0;
  color: transparent;
  background-color: rgba(0,0,0,.05);
  margin: 0 10px;
}
.uinfo-desc{
  padding: 10px 10px;
  overflow: hidden;
}
.uinfo-desc ul{
  list-style: none;
}
.uinfo-desc li{
  width: 50%;
  padding: 15px 0;
  text-align: center;
  display: inline-block;
  border-radius: 5px;
  cursor: pointer;
}
.uinfo-desc h3{
  color: #666;
  font-size: 13px;
  margin-bottom: 10px;
}
.uinfo-desc h1{
  color: #000;
  font-weight: 700;
  font-size: 16px;
}
.uinfo-desc li h1:hover{
  color: red;
}
.menulist ul{
  padding: 20px 8px;
  border: none;
  list-style: none;
}
.menu-item{
  padding: 0!important;
  text-align: center;
  height: 40px;
  line-height: 40px;
}

</style>