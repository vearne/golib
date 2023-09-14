package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"github.com/vearne/golib/metric"
	"time"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:3306",
		Password: "xxeQl*@nFE",
	})

	metric.AddRedis(redisClient, "test")

	for i := 0; i < 10; i++ {
		go func() {
			for {
				redisClient.Get(context.Background(), "a")
				time.Sleep(200 * time.Millisecond)
			}
		}()
	}

	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":9090")
}
