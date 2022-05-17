<template>
  <PersonalBar :title="personalBarTitle"></PersonalBar>
  <div class="content-main">
    <Profile></Profile>
    <router-view></router-view>
  </div>

</template>


<script>

import PersonalBar from "./PersonalBar";
import Profile from "./Profile";

export default {
  name: "PersonalCenter",
  data() {
    return {
      personalBarTitle: '个人中心'
    }
  },
  methods: {},
  components: {
    Profile,
    PersonalBar,
  },
}
</script>
<style scoped>
.content-main {
  display: flex;
  margin: 0 auto;
  padding: 130px 50px;
  background-color: #f7f7f7;
}
</style>

