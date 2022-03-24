<template>
  <Column></Column>
  <LoginPage v-if="true"></LoginPage>
  <RouteLine></RouteLine>
  <MainContext></MainContext>
</template>

<script>
import LoginPage from "@/components/entrance/LoginPage";
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
     Column, RouteLine, MainContext,LoginPage
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
