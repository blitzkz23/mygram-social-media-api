package entity

import "time"

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `gorm:"created_at,omitempty" json:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at,omitempty" json:"updated_at"`
}
