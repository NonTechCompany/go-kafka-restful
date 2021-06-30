package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func Init() {
	prometheus.Register(totalRequests)
}

func PrometheusMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		totalRequests.WithLabelValues(context.FullPath()).Inc()
	}
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)
