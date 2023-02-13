package handler

import (
	"context"
	"fmt"
	"os"

	"work/constants"
	"work/controller"

	"github.com/go-redis/redis/v8"
)

type Handler struct {
	redis *redis.Client
	cnt   *controller.Controller
}

//create new hander and conncet to redis
func NewHandler(cnt *controller.Controller) *Handler {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv(constants.RedisConnection),
	})
	
	return &Handler{redis: redisClient,
		cnt: cnt,
	}
}


//recive order from redis
func (h *Handler) SubOrder(ctx context.Context, chName string) {
	subscriber := h.redis.Subscribe(ctx, chName)

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")

		//send to save in database
		err = h.cnt.SaveInDatabase(msg.Payload)
		if err != nil {
			fmt.Println(err)
		}

	}

}
