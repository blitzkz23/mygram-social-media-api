package entity

import "time"

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `gorm:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at,omitempty"`
}
