import authRepository, {AuthRepository, IAccessTokenResponse, ISignupDTO} from "@repository/auth/repository";
import useAuthStore, {AuthUser} from "@store/auth/store";

export class AuthUsecase {
    private authRepository: AuthRepository = authRepository;

    async refreshAccessToken(): Promise<IAccessTokenResponse> {
        return this.authRepository.refreshAccessToken();
    }

    async signup(dto: ISignupDTO): Promise<void> {
        const data = await this.authRepository.signup(dto);
        useAuthStore().setAccessToken(data.accessToken)
    }

    async getAuthUser(): Promise<AuthUser> {
        return this.authRepository.getAuthUser();
    }
}

const authUsecase = new AuthUsecase();

export default authUsecase;