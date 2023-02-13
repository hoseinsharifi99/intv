package controller

import (
	"encoding/json"
	"fmt"
	"work/db_manager"
	"work/model"
)

type Controller struct {
	dm *db_manager.DbManager
}

func NewController(dm *db_manager.DbManager) *Controller {
	c := &Controller{dm: dm}
	return c
}

type authOrder struct {
	OrderID uint   `json:"order_id"`
	Price   uint   `json:"price"`
	Title   string `json:"title"`
}

//save order in database
func (c *Controller) SaveInDatabase(sOrder string) error {

	order := authOrder{}
	if err := json.Unmarshal([]byte(sOrder), &order); err != nil {
		panic(err)
	}

	dbOrder := &model.Order{
		OrderID: order.OrderID,
		Price:   order.Price,
		Title:   order.Title,
	}
	err := c.dm.AddOrder(dbOrder)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", order)
	return nil
}
