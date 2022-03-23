<template>
  <Column></Column>

  <RouteLine></RouteLine>
<!--  <Table :arr="array">
  </Table>-->
  <MainContext></MainContext>


</template>

<script>
// import Table from "@/components/Table";
import Column from "@/components/Bar";
import RouteLine from "@/components/RouteLine";
import MainContext from "@/components/MainContext";
export default {
  name: 'App',
  data() {
    return {
      array: [[1, 2, 3, 4], [1, 2, 3, 4]]
    }
  },
  components: {
     Column, RouteLine, MainContext
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

body {
  margin: 0 0;
}
</style>
