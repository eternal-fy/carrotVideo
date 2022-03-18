<template>
  <div id="column">
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
#column{
  position: fixed;
  width: 100%;
  padding: 20px 50px;
  z-index: 20;
  line-height: 50px;
  border-bottom: 1px solid black;
}
.context{
  display: flex;
  flex-direction: row;

}
</style>