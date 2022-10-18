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

func (c *commentPG) PostComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}
	comment.UserID = commentPayload.UserID

	if err := c.db.Model(&comment).Create(&commentPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	if err := c.db.Last(&comment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) GetAllComments() ([]*entity.Comment, errs.MessageErr) {
	comments := []*entity.Comment{}

	err := c.db.Preload("User").Preload("Photo").Find(&comments).Error
	if err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return comments, nil
}

func (c *commentPG) GetCommentByID(commentID uint) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}

	if err := c.db.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) EditCommentData(commentID uint, commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	comment := entity.Comment{}

	if err := c.db.Where("id = ?", commentID).Updates(&commentPayload).Error; err != nil {
		return nil, errs.NewInternalServerErrorr("Something went wrong")
	}

	return &comment, nil
}

func (c *commentPG) DeleteComment(commentID uint) errs.MessageErr {
	comment := entity.Comment{}

	if err := c.db.Where("id = ?", commentID).Delete(&comment).Error; err != nil {
		return errs.NewInternalServerErrorr("Something went wrong")
	}

	return nil
}
