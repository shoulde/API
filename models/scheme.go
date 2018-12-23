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
