package models

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" valid:"required-your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" valid:"required-your email is required, email-invalid format email"`
	Password string    `gorm:"not null" json:"password" valid:"password-your password is required, minstringlength(6)-password has to have a minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}
