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

const routerHistory = createWebHashHistory()
const router = createRouter({
    history: routerHistory,
    // scrollBehavior: true,
    routes: [

        {
            path: '/comedy',
            name: 'comedy',
            component: Comedy,
            meta: {
                requiresNotAuth: true
            }
        },
        {
            path: '/action',
            name: 'action',
            component: Action
        },
        {
            path: '/adventure',
            name: 'adventure',
            component: Adventure
        },
        {
            path: '/game',
            name: 'game',
            component: Game
        },
        {
            path: '/romance',
            name: 'romance',
            component: Romance
        },
        {
            path: '/thriller',
            name: 'thriller',
            component: Thriller
        },
        {
            path: '/video',
            name: 'video',
            component: Video
        },
        {
            path: '/vlog',
            name: 'vlog',
            component: Vlog
        },
        {
            path: '/war',
            name: 'war',
            component: War
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

    ]
})

export default router
