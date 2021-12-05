package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {

	router := gin.Default()
	router.GET("/visits", visits)
	router.Run(":8000")

}
func visits(c *gin.Context) {
	counter := 0
	fmt.Println("application started")
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val, err := rdb.Get(ctx, "counter").Result()
	if err == redis.Nil {
		counter = counter + 1
	} else {
		counter, _ = strconv.Atoi(val)
		counter = counter + 1
	}

	err = rdb.Set(ctx, "counter", counter, 0).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, counter)

}
