import {createApp} from 'vue'
import App from '@/Main.vue'
import router from '@/router'
import axios from 'axios'
import vueaxios from 'vue-axios'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'


const app = createApp(App);


app.config.globalProperties.$getCookie = (c_name) => {
    if (document.cookie.length > 0) {
        let c_start, c_end;
        c_start = document.cookie.indexOf(c_name + "=");
        if (c_start != -1) {
            c_start = c_start + c_name.length + 1;
            c_end = document.cookie.indexOf(";", c_start);
            if (c_end == -1) {
                c_end = document.cookie.length;
            }
            return document.cookie.substring(c_start, c_end);
        }
    }
    return "";
}
app.config.globalProperties.$deleteCookie = () => {
    document.cookie = "name=" + ""+";path=/";
/*    let now = new Date();
    now.setMonth(now.getMonth() - 1);
    document.cookie = "expires=" + now.toUTCString() + ";"+"path=/"*/
}

app.config.globalProperties.$checkCookie = (c_name) => {
    let username;
    username =app.config.globalProperties.$getCookie(c_name);
    if (username != null && username != "") {
        return true
    }
    return false
}

//默认使用axios
//更改axios默认路径，改为 /api 原因是要在vue.config.js里面配置
axios.defaults.baseURL = '/api'
app.config.globalProperties.$http = axios
app.use(vueaxios, axios)
app.use(router)
app.use(ElementPlus)
app.mount('#app')

