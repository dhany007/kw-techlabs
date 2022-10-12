package models

type Product struct {
	GormModel
	Title       string `json:"title" valid:"required-title of your product is required"`
	Description string `json:"description" valid:"required-description of your product is required"`
	UserID      uint
	User        *User
}
