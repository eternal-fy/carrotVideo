<template>
  <div id="square" >
      <router-link :to="context.link">{{context.name}}</router-link>
  </div>

</template>

<script>
export default {
  name: "square",
  props:{
    context:Object,

  }
}
</script>

<style scoped>
#square{
  margin: 8px 0 0 8px;
  font-size: 16px;
  width:120px;
  text-align:center ;
  font-family: Corbel;
  line-height: 40px;
}
#square:hover{
  background-color: beige;
}
a{
  text-decoration: none;
  font-weight: 500;
  color:#111622;
}
a:hover{
  color: aqua;
}

</style>
<style>

</style>