package main

import (
	"log"
	"time"

	"github.com/HenryNg101/golang-examples/sql_db/pkg"
	"github.com/HenryNg101/golang-examples/sql_db/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	redisClient := pkg.ConnectRedis()
	handler := &Handler{
		DBClient:    db.ConnectDB(),
		RedisClient: redisClient,
	}
	router := gin.Default()

	// Make use of Redis for rate limiting
	requestsLimitCount := 2 // Change value to something that you want. I used small value for testing purposes only
	router.Use(RateLimitMiddleware(redisClient, requestsLimitCount, time.Minute))

	router.GET("/users/top/:count", handler.TopSpendersHandler)

	// Cache invalidation demo. You can use get request first to see user's orders
	// Then, in the second request, updating the user's name would invalidate the key(s) related to the user
	router.GET("/users/:userid/orders", handler.GetUserOrdersHandler)
	router.PUT("/users/:userid", handler.UpdateUserHandler)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
