import {createRouter, createWebHistory} from 'vue-router'
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
import Face from "@/components/Face";
const RegisterPage = () => import("@/components/entrance/RegisterPage")
const PersonalCenter = () => import("@/components/personalInfo/PersonalCenter")
const VideoShow = () => import("@/components/personalInfo/VideoShow")
const PersonalVideos = () => import("@/components/personalInfo/PersonalVideos")
const PersonalInformation = () => import("@/components/personalInfo/PersonalInformation")
const routerHistory = createWebHistory()
const router = createRouter({
    history: routerHistory,
    // scrollBehavior: true,
    routes: [
        {
            path: '/',
            redirect: '/Face',
        },
        {
            path: '/Face',
            name: 'Face',
            component: Face,
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
        {
            path: '/personalCenter',
            name: 'personalCenter',
            component: PersonalCenter,
            meta: {
                requiresNotAuth: true
            },
            children: [
                {
                    path: 'personalVideos',
                    name: 'PersonalVideos',
                    component: PersonalVideos,
                },
                {
                    path: 'personalInformation',
                    name: 'PersonalInformation',
                    component: PersonalInformation,
                }
            ]
        },
        {
            path: '/videoShow/:id',
            name: 'videoShow',
            component: VideoShow,
            meta: {
                requiresNotAuth: true
            }
        },


    ]
})

export default router
