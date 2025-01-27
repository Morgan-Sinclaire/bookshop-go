package main

import (
	"github.com/morgan-sinclaire/bookshop-go/handlers"
	"github.com/gin-gonic/gin"
	"os"
	// "check"
	"log"
)

func main() {
	router := gin.Default()

	router.POST("/books/new", handlers.CreateBook)
	router.GET("/books/price", handlers.GetPrice)

	router.POST("/customers/new", handlers.CreateCustomer)
	router.PUT("/customers/updateAddress", handlers.UpdateCustomerAddress)
	router.GET("/customers/balance", handlers.GetCustomerBalance)

	router.POST("/orders/new", handlers.CreateOrder)
	router.GET("/orders/shipped", handlers.GetShipmentStatus)
	router.PUT("/orders/ship", handlers.ShipOrder)
	router.GET("/orders/status", handlers.GetOrderStatus)

	router.Run(":8080")

	file, _ := os.OpenFile("bookdrop.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// check(err)
	log.SetOutput(file)
}
