package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/HenryNg101/golang-examples/sql_db/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
	DBClient    *gorm.DB
	RedisClient *redis.Client
}

// Using cache-aside pattern, by checking if a key exist in cache first. If yes, use that
// If not, do normal query and then save in cache for future usage
func (handler *Handler) TopSpendersHandler(ctx *gin.Context) {
	count, err := strconv.Atoi(ctx.Param("count"))
	if err != nil || count <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid count"})
		return
	}

	// Check cache
	key := fmt.Sprintf("users:top:%d", count)
	cached, err := handler.RedisClient.Get(ctx, key).Result()
	if err == nil {
		var result []services.UserSpending
		err := json.Unmarshal([]byte(cached), &result)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"topUsers": result})
			return
		} else {
			log.Println("Error when loading JSON value from cache: ", err)
		}
	}

	// Cache miss, roll back to use DB
	result, err := services.GetTopSpenders(handler.DBClient, count)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save in cache for future use
	encodedResult, err := json.Marshal(result)
	// Add new key-value pair with TTL jitter to prevent stuff like cache stampede
	ttl := 5*time.Minute + time.Duration(rand.Intn(60))*time.Second
	handler.RedisClient.Set(ctx, key, encodedResult, ttl)

	ctx.JSON(http.StatusOK, gin.H{"topUsers": result})
}
