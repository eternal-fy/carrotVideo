<template>
  <div id="square" >
    {{context}}
  </div>

</template>

<script>
export default {
  name: "square",
  props:{
    context:String,

  }
}
</script>

<style scoped>
#square{
  width: 100%;
  height: 30px;
  font-size: 20px;
  align-content: center;
  font-family: Corbel
}
#square:hover{
  color: aqua;
}
</style>