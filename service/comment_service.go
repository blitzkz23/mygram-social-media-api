package service

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/repository/comment_repository"
)

type CommentService interface {
	PostComment(comment *entity.Comment) (*dto.CommentRequest, errs.MessageErr)
	GetAllComments() ([]*dto.CommentResponse, errs.MessageErr)
	GetCommentByID(commentID uint) (*dto.CommentResponse, errs.MessageErr)
	EditCommentData(commentID uint, comment *entity.Comment) (*dto.CommentRequest, errs.MessageErr)
	DeleteComment(commentID uint) errs.MessageErr
}

type commentService struct {
	commentRepository comment_repository.CommentRepository
}

func NewCommentService(commentRepository comment_repository.CommentRepository) CommentService {
	return &commentService{commentRepository: commentRepository}
}

func (c *commentService) PostComment(comment *entity.Comment) (*dto.CommentRequest, errs.MessageErr) {
	return nil, nil
}

func (c *commentService) GetAllComments() ([]*dto.CommentResponse, errs.MessageErr) {
	return nil, nil
}

func (c *commentService) GetCommentByID(commentID uint) (*dto.CommentResponse, errs.MessageErr) {
	return nil, nil
}

func (c *commentService) EditCommentData(commentID uint, comment *entity.Comment) (*dto.CommentRequest, errs.MessageErr) {
	return nil, nil
}

func (c *commentService) DeleteComment(commentID uint) errs.MessageErr {
	return nil
}
