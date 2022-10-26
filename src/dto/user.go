package dto

import "time"

type LoginRequest struct {
	Email    string `form:"email" json:"email" valid:"required~Email is required,email~Email is not valid" example:"naufaltampan@gmail.com"`
	Password string `form:"password" json:"password" valid:"required~Password is required" example:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hdWZhbHRhbXBhbkBnbWFpbC5jb20iLCJleHAiOjE2NjYyMjYwNjUsImlkIjozN30.Q0vWwNIom3ua1LpbyACM_zIIjXkq7AFN8U6YONL1lFM"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username" valid:"required~Username is required" example:"naufaltampan"`
	Email    string `json:"email" form:"email" valid:"required~Email is required,email~Email is not valid" example:"naufaltampan@gmail.com"`
	Password string `json:"password" form:"password" valid:"required~Password is required" example:"password"`
	Age      uint8  `json:"age" form:"age" valid:"required~Age is required,range(8|100)~Age must be between 8 and 100" example:"20"`
}

type RegisterResponse struct {
	ID       uint   `json:"id" `
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint8  `json:"age"`
}

type UpdateUserDataRequest struct {
	Email    string `json:"email" form:"email" valid:"required~Email is required,email~Email is not valid" example:"cubaganti@gmail.com"`
	Username string `json:"username" form:"username" valid:"required~Username is required" example:"cubaganti"`
}

type UpdateUserDataResponse struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       uint8      `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
