package handlers

import (
	"fmt"
	"net/http"

	"github.com/morgan-sinclaire/bookshop-go/db"
	"github.com/Morgan-Sinclaire/bookshop-go/logging"
	"github.com/gin-gonic/gin"
)

type Order struct {
	CustomerId int  `json:"customerId"`
	BookId     int  `json:"bookId"`
	Shipped    bool `json:"shipped"`
}

func CreateOrder(c *gin.Context) {
	var json Order
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if json.CustomerId <= 0 {
		c.JSON(400, gin.H{"error": "customerId must be a positive integer"})
		logging.LogMessage("Error: customerId must be a positive integer")
		return
	}

	if json.BookId <= 0 {
		c.JSON(400, gin.H{"error": "bookId must be a positive integer"})
		logging.LogMessage("Error: bookId must be a positive integer")
		return
	}

	_, err := db.CreatePO(json.BookId, json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success"})
}


func GetShipmentStatus(c *gin.Context) {
	var json Order
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pid, err := db.GetPOByContents(json.BookId, json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	status, err := db.IsPOShipped(pid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": status})
}

func ShipOrder(c *gin.Context) {
	var json Order
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pid, err := db.GetPOByContents(json.BookId, json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = db.ShipPO(pid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func GetOrderStatus(c *gin.Context) {
	var json Order
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pid, err := db.GetPOByContents(json.BookId, json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	addr, err := db.GetCustomerAddress(json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<html>
		<head>
		<title>Order Status</title>
		</head>
		<body>
		<h1>Order Status</h1>
		<p>Order ID: `+fmt.Sprint(pid)+`</p>
		<p>Book ID: `+fmt.Sprint(json.BookId)+`</p>
		<p>Customer ID: `+fmt.Sprint(json.CustomerId)+`</p>
		<p>Shipping Address: `+fmt.Sprint(addr)+`</p>
		</body>
	</html>`))
}
