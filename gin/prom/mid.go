package prom

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func StartProm() {
	router := gin.Default()

	// 创建和注册指标
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		// 定义标签（Label）用于区分指标，如请求方法、HTTP 状态码等
		[]string{"method", "status_code"},
	)
	prometheus.MustRegister(counter)

	// 将 Prometheus 中间件注册到 Gin 引擎中
	router.Use(PrometheusMiddleware())

	// 添加路由
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 启动 HTTP 服务器
	router.Run(":8080")
}

// PrometheusMiddleware 是将 promhttp.Handler() 注册为 Gin 中间件的函数
func PrometheusMiddleware() gin.HandlerFunc {
	handler := promhttp.Handler()

	return func(c *gin.Context) {
		if c.Request.URL.Path == "/metrics" {
			handler.ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}

		c.Next()
	}
}
