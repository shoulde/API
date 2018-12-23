package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"email,omitempty"`
	Email    string `json:"password,omitempty"`
	Password string `json:"name,omitempty"`
}

func (u *User) TableName() string {
	return "user"
}

type Photo struct {
	gorm.Model
	URL    string `json:"url,omitempty"`
	UserID int    `json:"userID,omitempty"`
}

func (p *Photo) TableName() string {
	return "photo"
}
