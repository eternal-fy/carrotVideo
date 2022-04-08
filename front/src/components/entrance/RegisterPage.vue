<template>
  <div class="registerPage">
    <RegisterForm></RegisterForm>
  </div>
</template>

<script>
import RegisterForm from "@/components/entrance/RegisterForm";
export default {
  name: "RegisterPage",
  components: {RegisterForm}
}
</script>

<style scoped>
.registerPage{
  width: 100%;
  height: 768px;
  background-size: cover;
}
</style>