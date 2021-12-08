package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {

	router := gin.Default()
	router.GET("/visits", visits)
	router.Run(":8001")

}
func visits(c *gin.Context) {
	counter := 0
	fmt.Println("application started")
	// ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	val, err := rdb.Get("counter").Result()
	if err == redis.Nil {
		counter = counter + 1
	} else {
		counter, _ = strconv.Atoi(val)
		counter = counter + 1
	}

	err = rdb.Set("counter", counter, 0).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, counter)

}
