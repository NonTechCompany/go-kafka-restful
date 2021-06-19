package user

import (
	"com.example.go-postgres/kafka_event"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func insertUser(user UserDTO) uint {

	entity := insert(User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})

	content, _ := json.Marshal(entity)
	kafka_event.EmployeeProducer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &kafka_event.EmployeeTopic, Partition: kafka.PartitionAny}, Value: content}

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
