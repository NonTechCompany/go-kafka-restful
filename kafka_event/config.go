package kafka_event

import "github.com/NonTechCompany/go-kafka-restful/config"

var employeeTopic string

func Init() {
	employeeTopic = config.GetAppConfig().Kafka.TopicName
	initProducer()
	initConsumer()
}
