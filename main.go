package main

import (
	"encoding/json"
	"github.com/NonTechCompany/go-kafka-restful/config"
	"github.com/NonTechCompany/go-kafka-restful/db"
	"github.com/NonTechCompany/go-kafka-restful/err"
	"github.com/NonTechCompany/go-kafka-restful/kafka_event"
	"github.com/NonTechCompany/go-kafka-restful/metric"
	"github.com/NonTechCompany/go-kafka-restful/user"
	"github.com/gin-gonic/gin"
	"time"
)

type CustomLog struct {
	TimeStamp  string `json:"time"`
	ClientIP   string
	Method     string
	Path       string
	Latency    float64
	StatusCode int
	Message    string
}

func main() {
	Init()
	r := gin.Default()
	r.Use(err.ErrorHandle())
	r.Use(metric.PrometheusMiddleware())
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {

		log, _ := json.Marshal(CustomLog{
			TimeStamp:  params.TimeStamp.Format(time.RFC1123),
			Latency:    params.Latency.Seconds(),
			ClientIP:   params.ClientIP,
			Method:     params.Method,
			Path:       params.Path,
			StatusCode: params.StatusCode,
			Message:    params.ErrorMessage,
		})
		return string(log) + "\n"
	}))

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
