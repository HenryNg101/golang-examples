package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/HenryNg101/golang-examples/sql_db/pkg/models"
	"github.com/HenryNg101/golang-examples/sql_db/pkg/services"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name"`
}

func (handler *Handler) UpdateUserHandler(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil || userId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user ID"})
		return
	}

	var req UserInfo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.UpdateUser(handler.DBClient, uint(userId), req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	keyPattern := fmt.Sprintf("users:%d*", userId)
	invalidKeys, err := handler.RedisClient.Keys(ctx, keyPattern).Result()
	if err == nil && len(invalidKeys) > 0 {
		removedCnt, err := handler.RedisClient.Del(ctx, invalidKeys...).Result()
		if err != nil || removedCnt < int64(len(invalidKeys)) {
			log.Println(err)
			log.Println("Not all keys are invalidated successfully")
		} else {
			log.Println("All keys are invalidated successfully")
		}
	}
}

func (handler *Handler) GetUserOrdersHandler(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil || userId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user ID"})
		return
	}

	// Check cache
	key := fmt.Sprintf("users:%d:orders", userId)
	cached, err := handler.RedisClient.Get(ctx, key).Result()
	var result models.User
	// var result map[string]interface{}
	if err == nil {
		err := json.Unmarshal([]byte(cached), &result)
		if err == nil {
			ctx.JSON(http.StatusOK, result)
			return
		} else {
			log.Println("Error when loading JSON value from cache: ", err)
		}
	}

	// Cache miss, roll back to use DB
	result, err = services.GetUserOrders(handler.DBClient, uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save in cache for future use
	encodedResult, err := json.Marshal(result)
	handler.RedisClient.Set(ctx, key, encodedResult, 5*time.Minute)

	ctx.JSON(http.StatusOK, result)
}
