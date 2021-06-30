package main

import (
	"github.com/NonTechCompany/go-kafka-restful/config"
	"github.com/NonTechCompany/go-kafka-restful/db"
	"github.com/NonTechCompany/go-kafka-restful/kafka_event"
	"github.com/NonTechCompany/go-kafka-restful/metric"
	"github.com/NonTechCompany/go-kafka-restful/user"
	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	r := gin.Default()
	r.Use(metric.PrometheusMiddleware())
	user.Router(r.Group("/user"))
	metric.Router(r.Group("/metrics"))
	_ = r.Run()
}

func Init() {
	config.Init()
	db.Init()
	metric.Init()
	_ = db.Connection.AutoMigrate(&user.User{})
	kafka_event.Init()

}
