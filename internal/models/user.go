package models

import (
	"gorm.io/gorm"
	"time"
)

type (
	UserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	UserResponse struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	User struct {
		gorm.Model
		Id        string `gorm:"primarykey"`
		Name      string
		Email     string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)
