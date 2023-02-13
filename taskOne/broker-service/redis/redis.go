package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	redis *redis.Client
}

var ctx = context.Background()

func NewConnetion() *Redis {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &Redis{redis: redisClient}
}

func (r *Redis) PublishOrder(order []byte, queueName string) error {

	if err := r.redis.Publish(ctx, "send-order-data", order).Err(); err != nil {
		panic(err)
	}
	return nil
}
