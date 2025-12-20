import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import useAuthStore from "@store/auth/store";
import {
    AdminPanelInfoView,
    AdminPanelView,
    AdminTablesView,
    DashboardView,
    LoginView,
    SignupView
} from "@app/router/views";
import {AuthPages, RouteName} from "@app/router/types";


const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: RouteName.DASHBOARD,
        component: DashboardView,
    },
    {
        path: '/signup',
        name: RouteName.SIGNUP,
        component: SignupView,
    },
    {
        path: '/login',
        name: RouteName.LOGIN,
        component: LoginView,
    },
    {
        path: '/admin',
        name: RouteName.ADMIN_PANEL,
        component: AdminPanelView,
        meta: { adminOnly: true },
        children: [
            {
                path: ':entity',
                children: [
                    {
                        path: '',
                        name: RouteName.ADMIN_PANEL_TABLE,
                        component: AdminTablesView,
                    },
                    {
                        path: 'info/:entityId',
                        name: RouteName.ADMIN_PANEL_INFO,
                        component: AdminPanelInfoView,
                    },
                    {
                        path: 'new',
                        name: RouteName.ADMIN_PANEL_NEW,
                        component: AdminPanelInfoView,
                    },
                ],
            },
        ],
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