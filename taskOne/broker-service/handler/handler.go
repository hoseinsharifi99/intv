package handler

import (
	"broker/controller"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	controller *controller.Controller
	ech        *echo.Echo
}

//creare newHander
func NewHandler(cnt *controller.Controller) *Handler {
	h := &Handler{controller: cnt, ech: echo.New()}
	h.defineRoutes()
	return h
}

func (h *Handler) defineRoutes() {
	h.ech.POST("/order", h.order)
}



func (h *Handler) Start(address string) {
	h.ech.Logger.Fatal(h.ech.Start(address))
}