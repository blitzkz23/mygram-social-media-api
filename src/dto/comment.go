package dto

import "time"

type CommentRequest struct {
	Message string `form:"message" json:"message" valid:"required~Message is required" example:"This is a comment"`
	PhotoID uint   `form:"photo_id" json:"photo_id" valid:"required~Photo ID is required" example:"1"`
}

type CommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type EmbeddedPhotoResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type GetCommentResponse struct {
	ID        uint                  `json:"id"`
	Message   string                `json:"message"`
	PhotoID   uint                  `json:"photo_id"`
	UserID    uint                  `json:"user_id"`
	CreatedAt *time.Time            `json:"created_at"`
	UpdatedAt *time.Time            `json:"updated_at"`
	User      EmbeddedUserResponse  `json:"User"`
	Photo     EmbeddedPhotoResponse `json:"Photo"`
}

type UpdateCommentRequest struct {
	Message string `form:"message" json:"message" valid:"required~Message is required" example:"This is a comment"`
}

type UpdateCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
