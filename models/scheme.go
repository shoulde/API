package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
	Photos   []Photo
}

func (u *User) TableName() string {
	return "user"
}

type Photo struct {
	gorm.Model
	URL    string
	UserID int `gorm:"index"`
}

func (p *Photo) TableName() string {
	return "photo"
}
