package kafka_event

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func InitConsumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{EmployeeTopic}, nil)

	go func() {
		for {
			fmt.Println("waiting for next message.....")
			msg, err := consumer.ReadMessage(-1)

			if err == nil {
				fmt.Printf("Message on %s: %s\n", *msg.TopicPartition.Topic, string(msg.Value))
				switch *msg.TopicPartition.Topic {
				case EmployeeTopic:
					fmt.Println("Some Employee Logic")
				}
			} else {
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
		consumer.Close()
	}()
}
