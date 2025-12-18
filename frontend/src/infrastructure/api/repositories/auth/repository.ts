import api, { Api } from "../../api";
import {AuthUser} from "@store/auth/store";

export interface ISignupDTO {
    username: string;
    email: string;
    password: string;
}

export interface ILoginDTO {
    usernameOrEmail: string;
    password: string;
}

export interface IAccessTokenResponse {
    accessToken: string;
}

export class AuthRepository {
    private readonly api: Api = api;

    async refreshAccessToken(): Promise<IAccessTokenResponse> {
        return this.api.get("/auth/refresh/access", { auth: false })
    }

    async signup(dto: ISignupDTO): Promise<IAccessTokenResponse> {
        return this.api.post("/auth/signup", dto, { auth: false })
    }

    async login(dto: ILoginDTO): Promise<IAccessTokenResponse> {
        return this.api.post("/auth/login", dto, { auth: false })
    }

    async getAuthUser(): Promise<AuthUser> {
        return this.api.get("/auth/user")
    }
}

const authRepository = new AuthRepository();

export default authRepository;