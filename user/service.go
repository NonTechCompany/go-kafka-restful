package user

import (
	"com.example.go-kafka-restful/kafka_event"
)

func insertUser(user UserDTO) uint {

	entity := insert(User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})

	kafka_event.Publish(entity)
	return entity.ID
}

func getUserFromId(id int) UserDTO {

	entity := findUserById(id)

	return UserDTO{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}
