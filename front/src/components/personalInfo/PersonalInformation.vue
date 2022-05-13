<template>
  <div class="p-information">

    <el-form :model="form" label-width="120px" :disabled="!infoEnable">
      <el-form-item label="name">
        <el-input v-model="form.name"/>
      </el-form-item>
      <el-form-item label="age">
        <el-input v-model.number="form.age" type="number" min="0" max="120"
                  autocomplete="off"/>
      </el-form-item>
      <el-form-item label="gender">
        <el-switch v-model="genderShow"/>
        {{ genderShow ? '男' : '女' }}
      </el-form-item>
      <el-form-item label="phone" prop="phone">
        <el-input v-model="form.phone" type="text" minlength="13" maxlength="13"
        />
      </el-form-item>
      <el-form-item label="address">
        <el-input v-model="form.address" type="textarea"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">submit</el-button>
        <el-button @click="this.infoEnable = false">Cancel</el-button>
      </el-form-item>
    </el-form>
    <el-button v-if="this.infoEnable==false" @click="this.infoEnable = true">change</el-button>
    <el-button @click="getCookies">getCookies</el-button>
  </div>
</template>

<script>
export default {
  name: "PersonalInformation",
  data() {
    return {
      genderShow: true,
      form: {
        name: '',
        age: Number,
        gender: 1,
        phone: '',
        address: '',
      },
      infoEnable: true
    }
  },
  methods: {
    onSubmit: function () {
      this.$http.post("user/saveinformation", {
        personalInfo: this.form
      }).then(res => {
        if (res.data.Code == "9999") {
          alert("修改成功！")
          this.infoEnable = false
          return
        }
        alert("修改失败！")
      })

    },
    getCookies: function () {
      console.log(document.cookie)
      console.log(this.$getCookie("name"))
      console.log(this.$getCookie("rand-sequence"))
    },

  },
  watch: {
    genderShow: {
      handler(newName) {
        if (newName) {
          this.form.gender = 1
        } else {
          this.form.gender = 0
        }
      }
    }
  }

}

</script>
<style scoped>
.p-information {
  padding-left: 100px;
}
</style>
