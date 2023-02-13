package controller

import (
	"broker/redis"
)

type Controller struct {
	redis *redis.Redis
}

func NewController(redis *redis.Redis) *Controller {
	c := &Controller{redis: redis}

	return c
}

// send order to redis
func (c *Controller) PublishMessage(order []byte) error {

	err := c.redis.PublishOrder(order, "orders")
	if err != nil {
		return err
	}

	return nil
}
