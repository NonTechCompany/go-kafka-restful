package kafka_event

var employeeTopic = "Employee"

func Init() {
	initProducer()
	initConsumer()
}
