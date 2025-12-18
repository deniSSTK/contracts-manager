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

export interface IAuthResponse {
    accessToken: string;
    exp: number;
}

export class AuthRepository {
    private readonly api: Api = api;

    async refreshAccessToken(): Promise<IAuthResponse> {
        return this.api.get("/auth/refresh/access", { auth: false })
    }

    async signup(dto: ISignupDTO): Promise<IAuthResponse> {
        return this.api.post("/auth/signup", dto, { auth: false })
    }

    async login(dto: ILoginDTO): Promise<IAuthResponse> {
        return await this.api.post("/auth/login", dto, { auth: false })
    }

    async getAuthUser(): Promise<AuthUser> {
        return this.api.get("/auth/user")
    }
}

const authRepository = new AuthRepository();

export default authRepository;