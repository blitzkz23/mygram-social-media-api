package comment_repository

import (
	"mygram-social-media-api/src/entity"
	"mygram-social-media-api/src/pkg/errs"
)

type CommentRepository interface {
	PostComment(comment *entity.Comment) (*entity.Comment, errs.MessageErr)
	GetAllComments() ([]*entity.Comment, errs.MessageErr)
	GetCommentByID(commentID uint) (*entity.Comment, errs.MessageErr)
	EditCommentData(commentID uint, comment *entity.Comment) (*entity.Comment, errs.MessageErr)
	DeleteComment(commentID uint) errs.MessageErr
}
