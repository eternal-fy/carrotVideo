<template>
  <PersonalBar></PersonalBar>
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
      uploadUrl: 'https://eternalfy.site/api/file/uploadfiles',
      localUrl: '/file/uploadfiles',
      file: null
    }
  },
  methods: {
  },
  components: {
    Profile,
    PersonalBar,
  },
}
</script>
<style scoped>
.content-main {
  display: flex;
  width: 100%;
  height: 500px;
  margin: 0 auto;
  padding: 130px 50px;
  background-color: #f7f7f7;
}
</style>

