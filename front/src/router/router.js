import Router from 'vue-router'
import Vue from "Vue"
import Table from "@/components/Table";
import RouteLine from "@/components/RouteLine";
Vue.use(Router)

const router = new Router({
    mode: 'history',
    // scrollBehavior: true,
    routes: [
        {
            path: '/table',
            name: 'table',
            component: Table
        },
        {
            path: '/routeline',
            name: 'routeline',
            component: RouteLine,
            meta:{
                requiresNotAuth: true
            }
        },
    ]
})
export default router
