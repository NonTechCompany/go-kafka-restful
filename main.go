package main

import (
	"com.example.go-postgres/db"
	"com.example.go-postgres/kafka_event"
	"com.example.go-postgres/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()
	kafka_event.Init()
	user.Router(r.Group("/user"))
	_ = r.Run()
}
