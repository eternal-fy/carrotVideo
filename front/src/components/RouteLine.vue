<template>
  <div id="routeLine">
    <Square v-for="(item,index) in routeLines.data" :key="index" :context="item">
    </Square>

    <div style="width:300px;margin:0 auto; padding:10px 0;" class="entranceIcon">
      <a target="_blank" href="http://beian.miit.gov.cn/" style="color:#666;margin-right:19px;">
        <p>鄂ICP备 2022005367号-1</p>
      </a>
      <div>
        <a target="_blank" href="http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=42011202002095"><img
            src="../assets/entranceIcon.png" style="float:left;"/>
          <p>鄂公网安备</p>
          <br/>
          <p>42011202002095号</p>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import routeLineConfig from "@/config/catalog.json"
import Square from "@/components/Square"

export default {
  name: "RouteLine",
  data() {
    return {
      routeLines: Object
    }
  },
  methods: {},
  mounted() {
    this.routeLines = routeLineConfig
  },
  components: {
    Square
  }
}
</script>

<style scoped>
#routeLine {
  background-color: #ffffff;
  width: 150px;
  padding: 20px 20px;
  position: fixed;
  z-index: 888;
  top: 90px;
  bottom: 0px;
  border-right: 1px solid rgba(0, 0, 0, 0.11);

}

#routeLine:hover {
  overflow-y: scroll;
}

.entranceIcon {
  display: inline-block;
  text-decoration: none;
  height: 20px;
  font-size: 10px;
  line-height: 20px;
}
.entranceIcon a{
  cursor: pointer;
}
.entranceIcon p {
  float: left;
  height: 20px;
  line-height: 20px;
  margin: 0px 0px 0px 5px;
  color: #939393;
}


</style>