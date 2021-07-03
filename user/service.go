package user

import (
	"github.com/NonTechCompany/go-kafka-restful/kafka_event"
)

func insertUser(user UserDTO) uint {

	entity := insert(User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})

	kafka_event.Publish(entity)
	return entity.ID
}

func getUserFromId(id int) (UserDTO, error) {

	entity, err := findUserById(id)
	if err != nil {
		return UserDTO{}, err
	}
	return UserDTO{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}, nil
}
