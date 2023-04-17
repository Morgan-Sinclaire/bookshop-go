package handlers

import (
	"github.com/morgan-sinclaire/bookshop-go/db"
	// "bookshop-go/logging"
	"github.com/gin-gonic/gin"
	"log"
)

type Customer struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	ShippingAddr   string  `json:"shippingAddr"`
	AccountBalance float32 `json:"accountBalance"`
}

func CreateCustomer(c *gin.Context) {
    var json Customer
    if err := c.BindJSON(&json); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if json.Name == "" {
        c.JSON(400, gin.H{"error": "name is required"})
				log.Println("Error: name is required")
        return
    }

		if json.ShippingAddr == "" {
				c.JSON(400, gin.H{"error": "shipping address is required"})
				log.Println("Error: shipping address is required")
				return
		}

    _, err := db.CreateCustomer(json.Name, json.ShippingAddr)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, gin.H{"status": "success"})
}

func UpdateCustomerAddress(c *gin.Context) {
	var json Customer
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if json.ShippingAddr == "" {
			c.JSON(400, gin.H{"error": "new shipping address is required"})
			log.Println("Error: new shipping address is required")
			return
	}

	err := db.UpdateCustomerAddress(json.Id, json.ShippingAddr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func GetCustomerBalance(c *gin.Context) {
	var json Customer
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	balance, err := db.CustomerBalance(json.Id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"balance": balance})
}
