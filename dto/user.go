package dto

type LoginRequest struct {
	Email    string `form:"email" json:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string `form:"password" json:"password" valid:"required~Password is required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Username string `form:"username" json:"username" valid:"required~Username is required"`
	Email    string `form:"email" json:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string `form:"password" json:"password" valid:"required~Password is required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}
