package kafka_event

import "com.example.go-kafka-restful/config"

var employeeTopic string

func Init() {
	employeeTopic = config.GetAppConfig().Kafka.TopicName
	initProducer()
	initConsumer()
}
