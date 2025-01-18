import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        name: 'start',
        path: '/',
        component: () => import('@/views/user/start.vue'),
    },
    {
        name: 'login',
        path: '/login',
        component: () => import('@/views/user/login.vue'),
    },
    {
        name: '404',
        path: '/:catchAll(.*)',
        component: () => import('@/views/404.vue') //  当路由不存在则到该路由
    },
    {
        name: 'register',
        path: '/register',
        component: () => import('@/views/user/register.vue') //  当路由不存在则到该路由
    },
    {
        name: 'retrieve',
        path: '/retrieve',
        component: () => import('@/views/user/retrieve.vue') //  当路由不存在则到该路由
    },
    {
        name: 'home',
        path: '/home',
        component: () => import('@/views/main/index.vue'),
    },
    {
        name: 'about',
        path: '/about',
        component: () => import('@/views/main/about.vue'),
    },
    {
        name: 'manage',
        path: '/manage',
        component: () => import('@/views/main/manage.vue'),
    },
    {
        name: 'fruit',
        path: '/fruit',
        component: () => import('@/views/main/fruit.vue'),
    },
    {
        name: 'buy-detail',
        path: '/buy-detail',
        component: () => import('@/views/main/buy-detail.vue'),
    },
    {
        name: 'order',
        path: '/order',
        component: () => import('@/views/main/order.vue'),
    },
    {
        name: 'main',
        path: '/main',
        component: () => import('@/views/admin/main.vue'),
    },
    {
        name: 'cart',
        path: '/cart',
        component: () => import('@/views/main/cart.vue'),
    },
    {
        name: 'user-manager',
        path: '/user-manager',
        component: () => import('@/views/admin/user-manager.vue'),
    },
    {
        name: 'kind-manager',
        path: '/kind-manager',
        component: () => import('@/views/admin/kind-manager.vue'),
    },
    {
        name: 'order-manager',
        path: '/order-manager',
        component: () => import('@/views/admin/order-manager.vue'),
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

// 全局导守卫
router.beforeEach((to, from, next) => {
    // to 满足条件的url放过
    if (to.path == "/register" || to.path == "/retrieve" || to.path == "/" || to.path == "/404"|| to.path == "/login") {
        next()  // 直接放过
    } else {
        // 校验用户的，不受管理员影响
        const auth_token = localStorage.getItem("auth_token")
        // front的需要权限的url
        if (auth_token == "" || auth_token == null || auth_token == '') {
            next("/")
        } else {
            next()
        }
    }
})

export default router