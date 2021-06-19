package main

import (
	"com.example.go-kafka-restful/config"
	"com.example.go-kafka-restful/db"
	"com.example.go-kafka-restful/kafka_event"
	"com.example.go-kafka-restful/user"
	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	r := gin.Default()
	user.Router(r.Group("/user"))
	_ = r.Run()
}

func Init() {
	config.Init()
	db.Init()
	kafka_event.Init()

}
