import useAuthStore from "@store/auth/store";

type Methods = "GET" | "POST" | "PUT" | "DELETE";

interface RequestOptions {
    extraHeaders?: Record<string, string>;
    auth?: boolean;
}

export class Api {
    private backendUrl: string = import.meta.env.VITE_BACKEND_URL + "/api";

    private async getAuthHeaders(): Promise<Record<string, string>> {
        const authStore = useAuthStore();

        if (!authStore.isAuthenticated || !authStore.accessToken) {
            await authStore.setAccessTokenRequest();

            if (authStore.accessToken) {
                return {
                    Authorization: `Bearer ${authStore.accessToken}`,
                };
            }
            return {};
        }

        return {
            Authorization: `Bearer ${authStore.accessToken}`,
        };
    }

    private async request<T>(
        endpoint: string,
        method: Methods,
        body?: any,
        options: RequestOptions = {}
    ): Promise<T> {
        const { extraHeaders = {}, auth = true } = options;

        const headers: Record<string, string> = {
            "Content-Type": "application/json",
            ...(auth ? await this.getAuthHeaders() : {}),
            ...extraHeaders,
        };

        const res = await fetch(`${this.backendUrl}${endpoint}`, {
            method,
            headers,
            body: body ? JSON.stringify(body) : undefined,
            credentials: "include",
        });

        if (auth && res.status === 401) {
            const authStore = useAuthStore();
            authStore.logout();
            throw new Error("Unauthorized, logged out");
        }

        if (!res.ok) {
            const data: { error: string } = await res.json();
            throw new Error(`HTTP ${res.status}: ${data.error}`);
        }

        const contentType = res.headers.get("content-type");
        if (contentType && contentType.includes("application/json")) {
            return (await res.json()) as T;
        }

        return {} as T;
    }

    get<T>(endpoint: string, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, "GET", undefined, options);
    }

    post<T>(endpoint: string, body: any, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, "POST", body, options);
    }

    put<T>(endpoint: string, body: any, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, "PUT", body, options);
    }

    delete<T>(endpoint: string, body?: any, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, "DELETE", body, options);
    }
}

const api = new Api();

export default api;