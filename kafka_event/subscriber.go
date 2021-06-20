package kafka_event

import (
	"fmt"
	"github.com/NonTechCompany/go-kafka-restful/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func initConsumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.GetAppConfig().Kafka.Bootstrap.Servers,
		"group.id":          config.GetAppConfig().Kafka.Consumer.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{employeeTopic}, nil)

	go subScribeTopics(consumer)
}

func subScribeTopics(consumer *kafka.Consumer) {
	for {
		fmt.Println("waiting for next message.....")
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Message on %s: %s\n", *msg.TopicPartition.Topic, string(msg.Value))
			switch *msg.TopicPartition.Topic {
			case employeeTopic:
				fmt.Println("Some Employee Logic")
			}
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	consumer.Close()
}
