import api, { Api } from "../../api";
import {AuthUser} from "@store/auth/store";
import {User} from "@model/user/model";
import {ListResult} from "../../dto";
import Contract from "@model/contract/entity";

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

    async get(id: string): Promise<User> {
        return this.api.get(`/auth/user/${id}`)
    }

    async getAuthUser(): Promise<AuthUser> {
        return this.api.get("/auth/user/me")
    }

    async login(dto: ILoginDTO): Promise<IAuthResponse> {
        return await this.api.post("/auth/login", dto, { auth: false })
    }

    async signup(dto: ISignupDTO): Promise<IAuthResponse> {
        return this.api.post("/auth/signup", dto, { auth: false })
    }

    async update(dto: any, id: string): Promise<boolean> {
        try {
            await this.api.put(`/auth/user/${id}`, dto);
            return true
        } catch {
            return false
        }
    }

    async list(filters: string): Promise<ListResult<User>> {
        return this.api.get("/auth/user/?" + filters)
    }

    async refreshAccessToken(): Promise<IAuthResponse> {
        return this.api.get("/auth/refresh/access", { auth: false })
    }
}

const authRepository = new AuthRepository();

export default authRepository;