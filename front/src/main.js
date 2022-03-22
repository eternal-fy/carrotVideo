import { createApp } from 'vue'
import App from '@/MainPage.vue'
import router from '@/router'
import axios from 'axios'
import vueaxios from 'vue-axios'

//更改axios默认路径，改为 /api 原因是要在vue.config.js里面配置
axios.default.baseURL='/api'
const app=createApp(App);
//默认使用axios
app.use(vueaxios,axios)
app.use(router)
app.mount('#app')