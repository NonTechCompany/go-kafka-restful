package user

import (
	"com.example.go-kafka-restful/db"
)

func insert(user User) User {
	database := db.Connection
	database.Create(&user)
	return user
}

func findUserById(id int) User {

	var user User
	db.Connection.First(&user, id)
	return user
}
