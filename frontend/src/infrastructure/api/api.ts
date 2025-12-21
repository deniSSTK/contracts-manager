import useAuthStore from "@store/auth/store";

type Methods = "GET" | "POST" | "PUT" | "DELETE";

type ResponseType = 'blob' | 'json'

interface RequestOptions {
    extraHeaders?: Record<string, string>;
    auth?: boolean;
    responseType?: ResponseType;

    isFormData?: boolean;
}

export class Api {
    private backendUrl: string = import.meta.env.VITE_BACKEND_URL + "/api";

    private async getAuthHeaders(): Promise<Record<string, string>> {
        const authStore = useAuthStore();
        const now = Math.floor(Date.now() / 1000);

        if (!authStore.accessToken || !authStore.exp || authStore.exp <= now) {
            await authStore.setAccessTokenRequest();
        }

        if (!authStore.accessToken) return {};

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
        const {
            extraHeaders = {},
            auth = true,
            responseType = 'json',
            isFormData = false,
        } = options;

        const headers: Record<string, string> = {
            ...(auth ? await this.getAuthHeaders() : {}),
            ...extraHeaders,
        };

        if (!isFormData && responseType === "json") {
            headers["Content-Type"] = "application/json";
        }

        const res = await fetch(`${this.backendUrl}${endpoint}`, {
            method,
            headers,
            body: body
                ? isFormData
                    ? body
                    : JSON.stringify(body)
                : undefined,
            credentials: "include",
        });

        if (!res.ok) {
            const data: { error: string } = await res.json();
            throw new Error(`HTTP ${res.status}: ${data.error}`);
        }

        if (responseType === 'blob') {
            return (await res.blob()) as T;
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