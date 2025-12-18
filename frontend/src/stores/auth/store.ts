import { defineStore } from 'pinia'
import authUsecase from "@usecase/auth/usecase";
import {IAuthResponse} from "@repository/auth/repository";

export interface AuthUser {
    id: string;
    type: UserType;
}

export type UserType = 'regular' | 'admin'

const useAuthStore = defineStore('auth', {
    state: () => ({
        accessToken: null as string | null,
        exp: null as number | null,
        authUser: null as AuthUser | null,
        loaded: false
    }),

    getters: {
        isAuthenticated: (state) => {
            if (!state.accessToken || !state.authUser || !state.exp) return false
            const now = Math.floor(Date.now() / 1000)
            return state.exp > now
        },
    },

    actions: {
        setAccessToken(tokenInfo: IAuthResponse): void {
            this.accessToken = tokenInfo.accessToken
            this.exp = tokenInfo.exp
        },

        async setAccessTokenRequest(): Promise<void> {
            const data = await authUsecase.refreshAccessToken()
            this.setAccessToken(data)
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

        async setAllWithToken(tokenInfo: IAuthResponse): Promise<void> {
            try {
                this.setAccessToken(tokenInfo)
                await this.setAuthUserRequest()
            } finally {
                this.loaded = true
            }
        },

        logout(): void {
            this.accessToken = null
            this.authUser = null
            this.exp = null
            this.loaded = false
        },
    },
})

export default useAuthStore;