import authRepository, {AuthRepository, IAuthResponse, ILoginDTO, ISignupDTO} from "@repository/auth/repository";
import useAuthStore, {AuthUser} from "@store/auth/store";

export class AuthUsecase {
    private authRepository: AuthRepository = authRepository;

    async refreshAccessToken(): Promise<IAuthResponse> {
        return this.authRepository.refreshAccessToken();
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

    async getAuthUser(): Promise<AuthUser> {
        return await this.authRepository.getAuthUser();
    }
}

const authUsecase = new AuthUsecase();

export default authUsecase;