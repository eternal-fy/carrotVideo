<template>
  <div id="personalbar">
    <div class="p-logoArea"><a class="p-logo" href="/"></a></div>
    <div class="context">
      <div class="columnContxt">

      </div>

    </div>
  </div>

</template>

<script>
export default {
  name: "PersonalBar"
}
</script>

<style scoped>
#personalbar {
  background-color: #ffffff;
  position: fixed;
  width: 100%;
  height: 90px;
  z-index: 888;
  line-height: 50px;
  border-bottom: 1px solid rgba(0, 0, 0, 0.11);
}

.context {
  padding: 20px 50px;
  float: right;
}

.columnContxt {
  height: 100%;
}

.p-logoArea {
  padding-left: 20px;
  height: 100%;
  width: 200px;
  float: left;
}

.p-logo {
  height: 100%;
  width: 100%;
  background-image: url("../../assets/carrot.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  display: block;
}

a {
  text-decoration: none;
}
</style>