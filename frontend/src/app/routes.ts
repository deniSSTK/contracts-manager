import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import SignupView from "@view/auth/SignupView.vue";
import useAuthStore from "@store/auth/store";
import DashboardView from "@view/dashboard/DashboardView.vue";

enum RouteName {
    LOGIN = "login",
    SIGNUP = "signup",

    DASHBOARD = 'dashboard'
}

const AuthPages: RouteName[] = [RouteName.LOGIN, RouteName.SIGNUP]

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        name: RouteName.DASHBOARD,
        component: DashboardView,
        meta: {
            requiresAuth: true,
        }
    },
    {
        path: '/signup',
        name: RouteName.SIGNUP,
        component: SignupView
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

    if (!authStore.isAuthenticated) {
        try {
            await authStore.loadAll()
        } catch (error) {
            return { name: RouteName.SIGNUP }
        }
    }

    if (!authStore.isAuthenticated && !AuthPages.includes(to.name as RouteName)) {
        return { name: RouteName.SIGNUP }
    }

    if (authStore.isAuthenticated && AuthPages.includes(to.name as RouteName)) {
        return { name: RouteName.DASHBOARD }
    }
})


export default router