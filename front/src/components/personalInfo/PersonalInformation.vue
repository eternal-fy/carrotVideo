<template>
  <div class="p-information">
    <el-form :model="userInfo" label-width="120px" :disabled="!infoEnable">
      <el-form-item label="name">
        <el-input v-model="userInfo.Name"/>
      </el-form-item>
      <el-form-item label="age">
        <el-input v-model.number="userInfo.Age" type="number" min="0" max="120"
                  autocomplete="off"/>
      </el-form-item>
      <el-form-item label="gender">
        <el-switch v-model="genderShow"/>
        {{ genderShow ? '男' : '女' }}
      </el-form-item>
      <el-form-item label="phone" prop="phone">
        <el-input v-model="userInfo.Phone" type="text" minlength="13" maxlength="13"
        />
      </el-form-item>
      <el-form-item label="address">
        <el-input v-model="userInfo.Address" type="textarea"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">submit</el-button>
        <el-button @click="this.infoEnable = false">Cancel</el-button>
      </el-form-item>
    </el-form>
    <el-button v-if="this.infoEnable==false" @click="this.infoEnable = true">change</el-button>
  </div>
</template>

<script>
export default {
  name: "PersonalInformation",
  data() {
    return {
      genderShow: true,
      infoEnable: false,
      username: '',
      userInfo: Object
    }
  },
  mounted() {
    this.username = this.$route.query.username
    let sendData = new FormData();// 上传文件的data参数
    sendData.append('username', this.username);
    this.$http.post("user/getinformation",sendData)
        .then((res) => {
          console.log(res)
          if (res.data.Code == 9999) {
            this.userInfo = res.data.TransData
          }
        })
  },
  methods:
      {
        onSubmit: function () {
          this.$http.post("user/saveinformation", {
            personalInfo: this.userInfo
          }).then(res => {
            if (res.data.Code == "9999") {
              alert("修改成功！")
              this.infoEnable = false
              return
            }
            alert("修改失败！")
          })
        }
        ,
      }
  ,
  watch: {
    genderShow: {
      handler(newName) {
        if (newName) {
          this.userInfo.Gender = 1
        } else {
          this.userInfo.Gender = 0
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
