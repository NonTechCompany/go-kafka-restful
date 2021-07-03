package user

import (
	"errors"
	"github.com/NonTechCompany/go-kafka-restful/db"
	"gorm.io/gorm"
)

var UserNotFoundError = errors.New("User not found")

func insert(user User) User {
	database := db.Connection
	database.Create(&user)
	return user
}

func findUserById(id int) (User, error) {

	var user User
	first := db.Connection.First(&user, id)
	if first.Error != nil {
		if first.Error == gorm.ErrRecordNotFound {
			return User{}, UserNotFoundError
		}
		return User{}, first.Error
	}
	return user, nil
}
