<template>
  <div id="loginPage">
    <div class="blackWrap">
    </div>
    <LoginForm></LoginForm>
  </div>

</template>

<script>
import LoginForm from "@/components/entrance/LoginForm";
export default {
  name: "LoginPage",
   components: {LoginForm}
}
</script>

<style scoped>
#loginPage {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
  z-index: 999;
}

.blackWrap {
  width: 100%;
  height: 100%;
  position: fixed;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
}
</style>