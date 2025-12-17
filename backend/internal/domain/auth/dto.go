package auth

type LoginDTO struct {
	UsernameOrEmail string `binding:"required,min=3" json:"usernameOrEmail"`
	Password        string `binding:"required,min=8" json:"password"`
}

type SignupDTO struct {
	Username string `binding:"required,min=3,max=50" json:"username"`
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=6" json:"password"`
}
