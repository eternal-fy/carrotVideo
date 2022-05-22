<template>
  <PersonalBar :title="personalBarTitle"></PersonalBar>
  <div class="content-main">
    <Profile :username="username"></Profile>
    <router-view></router-view>
  </div>

</template>


<script>

import PersonalBar from "./PersonalBar";
import Profile from "./Profile";

export default {
  name: "PersonalCenter",
  mounted() {
    this.username = this.$route.query.username
  },
  data() {
    return {
      personalBarTitle: '个人中心',
      username:''
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

