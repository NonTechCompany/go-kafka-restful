package kafka_event

var EmployeeTopic = "Employee"

func Init() {
	initProducer()
	InitConsumer()
}
