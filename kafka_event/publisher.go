package kafka_event

import "github.com/confluentinc/confluent-kafka-go/kafka"

var EmployeeProducer *kafka.Producer

func initProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})
	if err != nil {
		panic(err)
	}
	EmployeeProducer = p
}
