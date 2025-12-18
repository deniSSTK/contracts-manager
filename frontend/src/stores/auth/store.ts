import { defineStore } from 'pinia'
import authUsecase from "@usecase/auth/usecase";

export interface AuthUser {
    id: string;
    type: UserType;
}

export type UserType = 'regular' | 'admin'

const useAuthStore = defineStore('auth', {
    state: () => ({
        accessToken: null as string | null,
        authUser: null as AuthUser | null,
    }),

    getters: {
        isAuthenticated: (state) => !!state.accessToken && !!state.authUser,
    },

    actions: {
        setAccessToken(token: string): void {
            this.accessToken = token
        },

        async setAccessTokenRequest(): Promise<void> {
            const data = await authUsecase.refreshAccessToken()
            this.accessToken = data.accessToken
        },

        async setAuthUserRequest(): Promise<void> {
            this.authUser = await authUsecase.getAuthUser()
        },

        async loadAll(): Promise<void> {
            await this.setAccessTokenRequest()
            await this.setAuthUserRequest()
        },

        logout(): void {
            this.accessToken = null
            this.authUser = null
        },
    },
})

export default useAuthStore;