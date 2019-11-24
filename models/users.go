package models

import (
	"shade/config"
)

func AddUser(u *User) (err error) {
	if err = config.DB.Table("user").Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers(u *[]User) (err error) {
	if err = config.DB.Table("user").Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(u *User, id int) (err error) {
	if err := config.DB.Table("user").Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}
