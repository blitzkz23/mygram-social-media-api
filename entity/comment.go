package entity

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint
	Message string `gorm:"not null" form:"message" json:"message" valid:"required~Message is required"`
	User    *User
	Photo   *Photo
}
