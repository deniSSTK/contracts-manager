package auth

type LoginDTO struct {
	UsernameOrEmail string `binding:"required,min=5" json:"usernameOrEmail"`
	Password        string `binding:"required,min=8" json:"password"`
}

type SignupDTO struct {
	Username string `binding:"required,min=5,max=50" json:"username"`
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=8" json:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	Exp         int64  `json:"exp"`
}
