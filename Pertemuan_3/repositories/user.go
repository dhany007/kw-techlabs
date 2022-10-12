package repositories

import (
	"sesi3/models"
	"sesi3/params"
	"time"
)

func InsertUser(userParams *params.CreateUser) *models.User {
	return &models.User{
		Name:      userParams.Name,
		Email:     userParams.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
