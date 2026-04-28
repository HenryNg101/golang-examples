package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// Rate limit based on IP address
func RateLimitMiddleware(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Identify client (IP or userID)
		clientID := ctx.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", clientID)

		// Increment counter
		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "rate limiter error"})
			ctx.Abort()
			return
		}

		// Set expiration only on first hit
		if count == 1 {
			redisClient.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			ctx.Abort()
			return
		}

		ctx.Next() // Pass control to next handler
	}
}
