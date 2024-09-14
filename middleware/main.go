package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Logger 定义一个中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 给Context设置一个值
		c.Set("key1", "value1")
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
	}
}

func main() {
	r := gin.Default()
	// 作用于全局
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 作用于单个路由
	//r.GET("/benchmark", MyB)

}
