package kafka_event

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var employeeProducer *kafka.Producer

func initProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})
	if err != nil {
		panic(err)
	}
	employeeProducer = p

	go func() {
		for e := range employeeProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()
}

func Publish(message interface{}) {
	content, _ := json.Marshal(message)
	employeeProducer.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &employeeTopic, Partition: kafka.PartitionAny}, Value: content}
}
