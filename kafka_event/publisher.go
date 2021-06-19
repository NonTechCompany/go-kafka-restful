package kafka_event

import "github.com/confluentinc/confluent-kafka-go/kafka"

var EmployeeProducer *kafka.Producer

var EmployeeTopic = "Employee"
