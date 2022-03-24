<template>
  <div id="bar">
    <div class="context">
      <div class="loginArea">
        <a href="https://www.baidu.com" target="_blank">
          <div>
<!--            <el-link href="https://element.eleme.io" target="_blank">个人中心</el-link>-->
            <img width="50" height="50" src="../assets/personalCenter.svg"/>
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Column"
}
</script>

<style scoped>
#bar {
  background-color: #E4E7ED;
  opacity: 0.5;
  position: fixed;
  width: 100%;
  height: 50px;
  padding: 20px 50px;
  z-index: 10;
  line-height: 50px;
  overflow: hidden;
  border-bottom: 1px solid rgba(0, 0, 0, 0.11);
}

.context {
  display: flex;
  flex-direction: row;
}

.loginArea {
  width: 500px;
  line-height: 50px;
  font-max-size: 16px;


}
</style>