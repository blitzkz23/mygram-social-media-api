package service

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/pkg/helpers"
	"mygram-social-media-api/repository/comment_repository"
)

type CommentService interface {
	PostComment(userID uint, commentPayload *dto.CommentRequest) (*dto.CommentResponse, errs.MessageErr)
	GetAllComments() ([]*dto.CommentResponse, errs.MessageErr)
	GetCommentByID(commentID uint) (*dto.CommentResponse, errs.MessageErr)
	EditCommentData(commentID uint, commentPayload *dto.CommentRequest) (*dto.CommentResponse, errs.MessageErr)
	DeleteComment(commentID uint) errs.MessageErr
}

type commentService struct {
	commentRepository comment_repository.CommentRepository
}

func NewCommentService(commentRepository comment_repository.CommentRepository) CommentService {
	return &commentService{commentRepository: commentRepository}
}

func (c *commentService) PostComment(userID uint, commentPayload *dto.CommentRequest) (*dto.CommentResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(commentPayload); err != nil {
		return nil, err
	}

	payload := &entity.Comment{
		Message: commentPayload.Message,
		PhotoID: commentPayload.PhotoID,
		UserID:  userID,
	}

	comment, err := c.commentRepository.PostComment(payload)
	if err != nil {
		return nil, err
	}

	response := &dto.CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}

	return response, nil
}

func (c *commentService) GetAllComments() ([]*dto.CommentResponse, errs.MessageErr) {
	return nil, nil
}

func (c *commentService) GetCommentByID(commentID uint) (*dto.CommentResponse, errs.MessageErr) {
	return nil, nil
}

func (c *commentService) EditCommentData(commentID uint, commentPayload *dto.CommentRequest) (*dto.CommentResponse, errs.MessageErr) {
	return nil, nil
}

func (c *commentService) DeleteComment(commentID uint) errs.MessageErr {
	return nil
}
