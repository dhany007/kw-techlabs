package repositories

import (
	"assignment1/models"
	"assignment1/params"
	"time"
)

func InsertUser(req *params.CreateUser) models.User {
	newUser := models.User{
		Nama:      req.Nama,
		Email:     req.Email,
		Pekerjaan: req.Pekerjaan,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return newUser
}

func FindUserByNomorAbsen(users []models.User, absen int) (models.User, bool) {
	currentUser := models.User{}
	if (len(users) < absen) || (absen <= 0) {
		return currentUser, false
	}

	currentUser = users[absen-1]

	return currentUser, true
}
