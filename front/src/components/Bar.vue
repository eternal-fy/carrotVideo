<template>
  <div id="bar">
    <div class="context">
      <div v-for="(num,index) in 3" :key="index">{{num}}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Column"
}
</script>

<style scoped>
#bar{
  background-color: black;
  opacity: 0.5;
  position: fixed;
  width: 100%;
  padding: 20px 50px;
  z-index: 20;
  line-height: 50px;
  border-bottom: 1px solid rgba(0, 0, 0, 0.11);
}
.context{
  display: flex;
  flex-direction: row;

}
</style>