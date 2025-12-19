import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import useAuthStore from "@store/auth/store";

export enum RouteName {
    LOGIN = "login",
    SIGNUP = "signup",

    DASHBOARD = 'dashboard',

    ADMIN_PANEL = 'admin-panel',
    ADMIN_PANEL_TABLE = 'admin-panel-table',
    ADMIN_PANEL_INFO = 'admin-panel-info',
}

const AuthPages: RouteName[] = [RouteName.LOGIN, RouteName.SIGNUP]

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: RouteName.DASHBOARD,
        component: () => import("@view/dashboard/DashboardView.vue"),
    },
    {
        path: '/signup',
        name: RouteName.SIGNUP,
        component: () => import("@view/auth/SignupView.vue"),
    },
    {
        path: '/login',
        name: RouteName.LOGIN,
        component: () => import("@view/auth/LoginView.vue"),
    },
    {
        path: '/admin',
        name: RouteName.ADMIN_PANEL,
        component: () => import("@view/admin/AdminPanelView.vue"),
        meta: {
            adminOnly: true,
        },
        children: [
            {
                path: ':entity',
                children: [
                    {
                        path: '',
                        name: RouteName.ADMIN_PANEL_TABLE,
                        component: () => import("@view/admin/AdminTablesView.vue"),
                    },
                    {
                        path: 'info/:entityId',
                        name: RouteName.ADMIN_PANEL_INFO,
                        component: () => import("@view/admin/AdminPanelInfoVIew.vue")
                    }
                ]
            },
        ]
    },
    {
        path: '/:pathMatch(.*)*',
        redirect: { name: RouteName.DASHBOARD },
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach(async (to) => {
    const authStore = useAuthStore()

    if (!authStore.loaded) {
        try {
            await authStore.loadAll()
        } catch (error) {
            if (to.name != RouteName.SIGNUP) {
                return { name: RouteName.LOGIN }
            }
        }
    }

    if (authStore.accessToken && !authStore.isAuthenticated) {
        await authStore.setAccessTokenRequest()
    }

    if (authStore.isAuthenticated && AuthPages.includes(to.name as RouteName)) {
        return { name: RouteName.DASHBOARD }
    } else if (!authStore.isAuthenticated && !AuthPages.includes(to.name as RouteName)) {
        return { name: RouteName.LOGIN }
    }

    if (to.meta.adminOnly) {
        try {
            await authStore.setAuthUserRequest()
        } catch (error) {
            console.log(error)
        }
        if (!authStore.isAdmin) {
            return { name: RouteName.DASHBOARD }
        }
    }
})


export default router