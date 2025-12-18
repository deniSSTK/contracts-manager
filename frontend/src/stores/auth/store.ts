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
        loaded: false
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
            this.setAccessToken(data.accessToken)
        },

        async setAuthUserRequest(): Promise<void> {
            this.authUser = await authUsecase.getAuthUser()
        },

        async loadAll(): Promise<void> {
            try {
                await this.setAccessTokenRequest()
                await this.setAuthUserRequest()
            } finally {
                this.loaded = true
            }
        },

        async setAllWithToken(token: string): Promise<void> {
            try {
                this.setAccessToken(token)
                await this.setAuthUserRequest()
            } finally {
                this.loaded = true
            }
        },

        logout(): void {
            this.accessToken = null
            this.authUser = null
        },
    },
})

export default useAuthStore;