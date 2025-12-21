import authRepository, {AuthRepository, IAuthResponse, ILoginDTO, ISignupDTO} from "@repository/auth/repository";
import useAuthStore, {AuthUser} from "@store/auth/store";
import {User} from "@model/user/model";
import {ListResult} from "../../infrastructure/api/dto";
import Contract from "@model/contract/entity";

export class AuthUsecase {
    private authRepository: AuthRepository = authRepository;

    async get(id: string): Promise<User> {
        return await this.authRepository.get(id)
    }

    async getAuthUser(): Promise<AuthUser> {
        return await this.authRepository.getAuthUser();
    }

    async login(dto: ILoginDTO): Promise<boolean> {
        const authStore = useAuthStore()

        try {
            const data = await this.authRepository.login(dto)
            await authStore.setAllWithToken(data)
        } catch {
            return false
        }
        return true
    }

    async signup(dto: ISignupDTO): Promise<boolean> {
        const authStore = useAuthStore()

        try {
            const data = await this.authRepository.signup(dto)
            await authStore.setAllWithToken(data)
        } catch {
            return false
        }
        return true
    }

    async update(dto: any, id: string): Promise<boolean> {
        return this.authRepository.update(dto, id);
    }

    async list(filters: string): Promise<ListResult<User>> {
        return await this.authRepository.list(filters)
    }

    async refreshAccessToken(): Promise<IAuthResponse> {
        return this.authRepository.refreshAccessToken();
    }

    async getUserContracts(): Promise<Contract[]> {
        return this.authRepository.getUserContracts()
    }

    async logout(): Promise<void> {
        return this.authRepository.logout()
    }
}

const authUsecase = new AuthUsecase();

export default authUsecase;