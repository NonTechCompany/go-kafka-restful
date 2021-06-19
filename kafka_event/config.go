package kafka_event

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Init() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})
	if err != nil {
		panic(err)
	}
	EmployeeProducer = p

	go Subscribe()
}
