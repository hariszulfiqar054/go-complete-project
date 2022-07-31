package services

import (
	"github.com/haris/go-rest/src/models"
	"github.com/haris/go-rest/src/repository"
)

func CreateAdmin(user *models.User) {
	repository.Create_User(user)
}
