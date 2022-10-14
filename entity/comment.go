package entity

type Comment struct {
	GormModel
	UserID  int64
	PhotoID int64
	Message string `gorm:"not null" form:"message" json:"message" valid:"required~Message is required"`
}
