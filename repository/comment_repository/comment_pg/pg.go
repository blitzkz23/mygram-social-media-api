package comment_pg

import (
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/repository/comment_repository"

	"gorm.io/gorm"
)

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) comment_repository.CommentRepository {
	return &commentPG{db: db}
}

func (c *commentPG) PostComment(comment *entity.Comment) (*entity.Comment, errs.MessageErr) {
	return nil, nil
}

func (c *commentPG) GetAllComments() ([]*entity.Comment, errs.MessageErr) {
	return nil, nil
}

func (c *commentPG) GetCommentByID(commentID uint) (*entity.Comment, errs.MessageErr) {
	return nil, nil
}

func (c *commentPG) EditCommentData(commentID uint, comment *entity.Comment) (*entity.Comment, errs.MessageErr) {
	return nil, nil
}

func (c *commentPG) DeleteComment(commentID uint) errs.MessageErr {
	return nil
}
