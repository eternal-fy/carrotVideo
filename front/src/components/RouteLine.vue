<template>
  <div id="routeLine">
    <Square v-for="(item,index) in 50" :key="index" :context="item" @click="click">
    </Square>
  </div>
</template>

<script>
import routeLineConfig from "@/config/routeLineConfig.json"
import Square from "@/components/Square";
export default {
  name: "RouteLine",
  data(){
    return {
      routeLines:{}
    }
  },
  methods:{
    click:function (c){
      console.log(c)
      console.log(this)
    }
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
  width: 150px;
  padding-top:20px;
  position: fixed;
  z-index: 10;
  overflow-y: scroll;
  top:90px;
  bottom: 0px;
  border-right: 1px solid black;

}
</style>