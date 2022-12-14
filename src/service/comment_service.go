package service

import (
	"mygram-social-media-api/src/dto"
	"mygram-social-media-api/src/entity"
	"mygram-social-media-api/src/pkg/errs"
	"mygram-social-media-api/src/pkg/helpers"
	"mygram-social-media-api/src/repository/comment_repository"
	"mygram-social-media-api/src/repository/photo_repository"
)

type CommentService interface {
	PostComment(userID uint, commentPayload *dto.CommentRequest) (*dto.CommentResponse, errs.MessageErr)
	GetAllComments() ([]*dto.GetCommentResponse, errs.MessageErr)
	EditCommentData(commentID uint, commentPayload *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr)
	DeleteComment(commentID uint) (*dto.DeleteCommentResponse, errs.MessageErr)
}

type commentService struct {
	commentRepository comment_repository.CommentRepository
	photoRepository   photo_repository.PhotoRepository
}

func NewCommentService(commentRepository comment_repository.CommentRepository, photoRepository photo_repository.PhotoRepository) CommentService {
	return &commentService{commentRepository: commentRepository, photoRepository: photoRepository}
}

func (c *commentService) PostComment(userID uint, commentPayload *dto.CommentRequest) (*dto.CommentResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(commentPayload); err != nil {
		return nil, err
	}

	entity := &entity.Comment{
		Message: commentPayload.Message,
		PhotoID: commentPayload.PhotoID,
		UserID:  userID,
	}

	_, err := c.photoRepository.GetPhotoByID(commentPayload.PhotoID)
	if err != nil {
		return nil, err
	}

	comment, err := c.commentRepository.PostComment(entity)
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

func (c *commentService) GetAllComments() ([]*dto.GetCommentResponse, errs.MessageErr) {
	comments, err := c.commentRepository.GetAllComments()
	if err != nil {
		return nil, err
	}

	response := make([]*dto.GetCommentResponse, 0)
	for _, comment := range comments {
		response = append(response, comment.ToGetCommentResponseDTO())
	}

	return response, nil
}

func (c *commentService) EditCommentData(commentID uint, commentPayload *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(commentPayload)
	if err != nil {
		return nil, err
	}

	entity := &entity.Comment{
		Message: commentPayload.Message,
	}

	comment, err := c.commentRepository.EditCommentData(commentID, entity)
	if err != nil {
		return nil, err
	}

	response := &dto.UpdateCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}

	return response, nil
}

func (c *commentService) DeleteComment(commentID uint) (*dto.DeleteCommentResponse, errs.MessageErr) {
	if err := c.commentRepository.DeleteComment(commentID); err != nil {
		return nil, err
	}

	response := &dto.DeleteCommentResponse{
		Message: "Your comment has been deleted",
	}

	return response, nil
}
