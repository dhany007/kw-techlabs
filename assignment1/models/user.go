package models

import (
	"fmt"
	"time"
)

type User struct {
	Nama      string
	Email     string
	Pekerjaan string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) PrintUser() {
	fmt.Println("Nama \t\t: ", user.Nama)
	fmt.Println("Email \t\t: ", user.Email)
	fmt.Println("Pekerjaan \t: ", user.Pekerjaan)
	fmt.Println("CreatedAt \t: ", user.CreatedAt)
	fmt.Println("UpdatedAt \t: ", user.UpdatedAt)
}
