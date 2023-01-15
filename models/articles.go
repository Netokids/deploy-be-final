package models

import "time"

type Article struct {
	ID          int          `json:"id" gorm:"primary_key:auto_increment"`
	Title       string       `json:"title" gorm:"type : varchar (255)"`
	Image       string       `json:"image" gorm:"type : varchar (255)"`
	Description string       `json:"description" gorm:"type : varchar (255)"`
	User        UserResponse `json:"user"`
	UserID      int          `json:"-"`
	CreatedAt   time.Time    `json:"-"`
}
