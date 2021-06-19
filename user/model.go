package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
}

type UserDTO struct {
	ID        uint
	FirstName string
	LastName  string
}
