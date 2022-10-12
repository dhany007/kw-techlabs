package models

import (
	"fmt"
	"time"
)

type User struct {
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) PrintUser() {
	fmt.Printf("Name: %s, Email: %s, CeatedAt: %s, UpdatedAt: %s", u.Name, u.Email, u.CreatedAt, u.UpdatedAt)
}
