package entity

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null;type:varchar(191)" form:"name" json:"name" valid:"required~ame is required"`
	SocialMediaURL string `gorm:"not null" form:"social_media_url" json:"social_media_url" valid:"required~Social Media URL is required"`
	UserID         uint
	User           *User
}
