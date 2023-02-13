package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authOrder struct {
	OrderID uint   `json:"order_id"`
	Price   uint   `json:"price"`
	Title   string `json:"title"`
}



//function for order endpoint 
func (h *Handler) order(c echo.Context) error {
	authOrderRequest := &authOrder{}
	if err := c.Bind(authOrderRequest); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "error binding user request")
	}

	orderBytes, err := json.Marshal(authOrderRequest)

	if err != nil {
		return err
	}

	//call publishMessage function to send order to redis publisher
	err = h.controller.PublishMessage(orderBytes)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "order added to queue")
}
