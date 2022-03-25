<template>
  <div id="loginPage" v-if="loginSwitch">
    <div class="blackWrap" >
      <el-icon class="closeIcon" @click="$emit('showSwitch')">
        <close-bold/>
      </el-icon>
    </div>
    <LoginForm @loginComplete="$emit('showSwitch')"></LoginForm>
  </div>

</template>

<script>
import LoginForm from "@/components/entrance/LoginForm";
import {CloseBold} from "@element-plus/icons-vue";

export default {
  name: "LoginPage",
  props: {
    loginSwitch: Boolean
  },
  methods: function loSwitch() {
  this.$emit("")
  },

  components: {LoginForm, CloseBold}
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

.closeIcon {
  position: absolute;
  right: 20px;
  color: white;
  font-size: 30px;

}
</style>