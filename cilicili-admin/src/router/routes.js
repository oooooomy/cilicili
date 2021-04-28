import Layout from '../layout/Index'

const routes = [
    {
        path: '/',
        component: Layout,
        children: [
            {
                path: "",
                redirect: "/system",
                meta: {auth: true}
            },
            {
                path: 'system',
                component: () => import('../views/system/Index'),
                meta: {auth: true}
            },
            {
                path: 'music',
                component: () => import('../views/music/MusicList'),
                meta: {auth: true}
            },
            {
                path: 'add',
                component: () => import('../views/music/AddMusic'),
                meta: {auth: true}
            },
            {
                path: 'audit',
                component: () => import('../views/room/Audit'),
                meta: {auth: true}
            },
            {
                path: 'live',
                component: () => import('../views/room/Live'),
                meta: {auth: true}
            },
            {
                path: 'worker',
                component: () => import('../views/setting/Worker'),
                meta: {auth: true}
            },
            {
                path: 'account',
                component: () => import('../views/setting/Account'),
                meta: {auth: true}
            },
            {
                path: 'videos',
                component: () => import('../views/video/Videos'),
                meta: {auth: true}
            },
            {
                path: 'wait',
                component: () => import('../views/video/Wait'),
                meta: {auth: true}
            },
            //403
            {
                path: '403',
                component: () => import('../views/error/403'),
                meta: {auth: true}
            },
            //404
            {
                path: '404',
                component: () => import('../views/error/404'),
                meta: {auth: true}
            },
            //500
            {
                path: '500',
                component: () => import('../views/error/500'),
                meta: {auth: true}
            },
        ]
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/Login'),
        meta: {auth: false}
    }
]

export default routes