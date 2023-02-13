package main

import (
	"broker/controller"
	"broker/redis"

	"broker/handler"
)

func main() {
	//create connection to redis
	redisConn := redis.NewConnetion()

	cnt := controller.NewController(redisConn)
	//create handler
	h := handler.NewHandler(cnt)

	//start service
	h.Start(":8080")
}
