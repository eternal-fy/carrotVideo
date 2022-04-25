<template>
  <div id="routeLine">
    <Square v-for="(item,index) in routeLines.data" :key="index" :context="item" >
    </Square>
  </div>
</template>

<script>
import routeLineConfig from "@/config/catalog.json"
import Square from "@/components/Square"

export default {
  name: "RouteLine",
  data(){

  },
  methods:{
  },
  created() {
    this.routeLines=routeLineConfig
  },
  components:{
    Square
  }
}
</script>

<style scoped>
#routeLine{
  background-color: #ffffff;
  width: 150px;
  padding:20px 20px;
  position: fixed;
  z-index: 888;
  top:90px;
  bottom: 0px;
  border-right: 1px solid rgba(0, 0, 0, 0.11);

}
#routeLine:hover{
  overflow-y: scroll;
}
</style>