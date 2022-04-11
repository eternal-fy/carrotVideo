import { createApp } from 'vue'
import App from '@/Main.vue'
import router from '@/router'
import axios from 'axios'
import vueaxios from 'vue-axios'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'


const app=createApp(App);
//默认使用axios
//更改axios默认路径，改为 /api 原因是要在vue.config.js里面配置
axios.defaults.baseURL='http://eternalfy.site:9999'
app.config.globalProperties.$http = axios
app.use(vueaxios,axios)
app.use(router)
app.use(ElementPlus)
app.mount('#app')

