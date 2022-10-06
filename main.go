package main

import (
	"log"

	"example.com/m/v2/db"
	"example.com/m/v2/myroute"
	"example.com/m/v2/redis"

	"github.com/gin-gonic/gin"
)

func main() {

	c := gin.Default()

	err := redis.InitRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	myroute.Routing(c)

	defer redis.CloseRedis()

	c.Run()
}
