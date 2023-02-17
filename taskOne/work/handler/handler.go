package handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"work/constants"
	"work/controller"

	"github.com/go-redis/redis/v8"
)

type Handler struct {
	redis *redis.Client
	cnt   *controller.Controller
}

// create new hander and conncet to redis
func NewHandler(cnt *controller.Controller) *Handler {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv(constants.RedisConnection),
	})

	return &Handler{redis: redisClient,
		cnt: cnt,
	}
}

// recive order from redis
func (h *Handler) SubOrder(ctx context.Context, chName string) {

	for {
		// Read the first item from the Redis list
		res, err := h.redis.RPop(ctx, chName).Result()
		if err == redis.Nil {
			// If the Redis list is empty, wait for 1 second before checking again
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			log.Println("Failed to retrieve order from queue:", err)
			continue
		}

		fmt.Println("Received message from " + chName + " channel.")
		fmt.Println(res)
		//send to save in database
		err = h.cnt.SaveInDatabase(res)
		if err != nil {
			fmt.Println(err)
		}

	}

}
