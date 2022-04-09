<template>
  <div>
    <router-link to="/main-page/main-context/index" target="_blank">
      <video class="page-media" autoplay muted loop name="medio">
        <source src="../assets/star.mp4" type="video/mp4"/>
      </video>
    </router-link>
  </div>
</template>

<script>
export default {
  name: "Face"
}
</script>

<style scoped>
.page-media {
  position: fixed;
  width: 100%;
  top: 0;
}

</style>