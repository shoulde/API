package models

import (
	"shade/config"
)

func AddUser(u *User) (err error) {
	if err = config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers(u *[]User) (err error) {
	if err = config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(u *User, id int) (err error) {
	if err := config.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}
