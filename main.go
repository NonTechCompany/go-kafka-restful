package main

import (
	"com.example.go-kafka-restful/db"
	"com.example.go-kafka-restful/kafka_event"
	"com.example.go-kafka-restful/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()
	kafka_event.Init()
	user.Router(r.Group("/user"))
	_ = r.Run()
}
