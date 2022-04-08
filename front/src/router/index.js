import {createRouter, createWebHashHistory} from 'vue-router'
import Table from "@/components/Table";
import RouteLine from "@/components/RouteLine";
import Action from "@/components/catagory/Action";
import Comedy from "@/components/catagory/Comedy";
import Adventure from "@/components/catagory/Adventure";
import Game from "@/components/catagory/Game";
import Romance from "@/components/catagory/Romance";
import Thriller from "@/components/catagory/Thriller";
import Video from "@/components/catagory/Video";
import Vlog from "@/components/catagory/Vlog";
import War from "@/components/catagory/War";
import mainContext from "@/components/MainContext";
import Index from "@/components/catagory/Index";
import mainPage from "@/MainPage";

const RegisterPage = () => import("@/components/entrance/RegisterPage")


const routerHistory = createWebHashHistory()
const router = createRouter({
    history: routerHistory,
    // scrollBehavior: true,
    routes: [
        {
            path: '/',
            redirect: '/main-page/main-context/index',
        },
        {
            path: '/main-page',
            component: mainPage,
            children: [{
                path: 'main-context',
                name: 'mainContext',
                component: mainContext,
                children: [
                    {
                        path: 'index',
                        name: 'index',
                        component: Index,
                        meta: {
                            requiresNotAuth: true
                        }
                    },
                    {
                        path: 'comedy',
                        name: 'comedy',
                        component: Comedy,
                        meta: {
                            requiresNotAuth: true
                        }
                    },
                    {
                        path: 'action',
                        name: 'action',
                        component: Action
                    },
                    {
                        path: 'adventure',
                        name: 'adventure',
                        component: Adventure
                    },
                    {
                        path: 'game',
                        name: 'game',
                        component: Game
                    },
                    {
                        path: 'romance',
                        name: 'romance',
                        component: Romance
                    },
                    {
                        path: 'thriller',
                        name: 'thriller',
                        component: Thriller
                    },
                    {
                        path: 'video',
                        name: 'video',
                        component: Video
                    },
                    {
                        path: 'vlog',
                        name: 'vlog',
                        component: Vlog
                    },
                    {
                        path: 'war',
                        name: 'war',
                        component: War
                    }
                ]
            }]
        },
        {
            path: '/table',
            name: 'table',
            component: Table
        },
        {
            path: '/routeline',
            name: 'routeline',
            component: RouteLine,
            meta: {
                requiresNotAuth: true
            }
        },
        {
            path: '/register',
            name: 'register',
            component: RegisterPage,
            meta: {
                requiresNotAuth: true
            }
        },


    ]
})

export default router
