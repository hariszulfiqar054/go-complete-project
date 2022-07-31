package repository

import (
	"github.com/haris/go-rest/src/models"
	"github.com/haris/go-rest/src/utils"
)

func Create_User(user *models.User) {
	result := utils.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
}
