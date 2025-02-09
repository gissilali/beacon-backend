package dtos

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	UserCredentials
	Name string `json:"name"`
}

type LoginUserRequest struct {
	UserCredentials
}

type AuthTokens struct {
	AccessToken  string `json:"-"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshToken struct {
	UserID int
}

type RefreshTokenRequest struct {
	RefreshToken string
}
