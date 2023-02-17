package main

import (
	"context"
	"fmt"
	"log"
	"work/controller"
	"work/db"
	"work/db_manager"
	"work/handler"

	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {
	//load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connect to database
	db := db.Connect()
	dm := db_manager.NewDbManager(db)

	cnt := controller.NewController(dm)
	h := handler.NewHandler(cnt)

	fmt.Println("=====SERVICE STARTED=====")

	h.SubOrder(ctx, "orders")
}
