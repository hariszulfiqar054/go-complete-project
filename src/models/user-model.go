package models

import (
	"time"

	"github.com/haris/go-rest/src/constants"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Password  string         `js1on:"password"`
	Role      constants.Role `json:"role" sql:"type:ENUM('admin','user')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
