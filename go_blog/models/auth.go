package models

import (
	"github.com/jinzhu/gorm"
)

type Blog_auth struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//CheckAuth checks if authentication information exists
func CheckAuth(username string, password string) (bool, error) {
	var auth Blog_auth
	err := db.Select("id").Where(Blog_auth{
		Username: username,
		Password: password}).First(&auth).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
