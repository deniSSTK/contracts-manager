import authRepository, {AuthRepository, IAccessTokenResponse, ILoginDTO, ISignupDTO} from "@repository/auth/repository";
import useAuthStore, {AuthUser} from "@store/auth/store";

export class AuthUsecase {
    private authRepository: AuthRepository = authRepository;

    async refreshAccessToken(): Promise<IAccessTokenResponse> {
        return this.authRepository.refreshAccessToken();
    }

    async signup(dto: ISignupDTO): Promise<boolean> {
        const authStore = useAuthStore()

        try {
            const data = await this.authRepository.signup(dto);
            await authStore.setAllWithToken(data.accessToken)
        } catch {
            return false
        }
        return true
    }

    async login(dto: ILoginDTO): Promise<boolean> {
        const authStore = useAuthStore()

        try {
            const data = await this.authRepository.login(dto)
            await authStore.setAllWithToken(data.accessToken)
        } catch {
            return false
        }
        return true
    }

    async getAuthUser(): Promise<AuthUser> {
        return this.authRepository.getAuthUser();
    }
}

const authUsecase = new AuthUsecase();

export default authUsecase;