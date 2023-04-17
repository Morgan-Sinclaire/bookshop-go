package handlers

import (
	"github.com/morgan-sinclaire/bookshop-go/db"
	"github.com/Morgan-Sinclaire/bookshop-go/logging"
	"github.com/gin-gonic/gin"
	// "log"
)

type Book struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func CreateBook(c *gin.Context) {
	var json Book
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if json.Title == "" {
		c.JSON(400, gin.H{"error": "title field is required"})
		logging.LogMessage("Error: title is required")
		return
	}

	if json.Author == "" {
		c.JSON(400, gin.H{"error": "author field is required"})
		logging.LogMessage("Error: author is required")
		return
	}

	if json.Price <= 0 {
		c.JSON(400, gin.H{"error": "positive price is required"})
		logging.LogMessage("Error: positive price is required")
		return
	}

	_, err := db.CreateBook(json.Title, json.Author, json.Price)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success"})
}

func GetPrice(c *gin.Context) {
	var json Book
	if err := c.BindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	bid, err := db.GetBookId(json.Title, json.Author)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	price, err := db.GetBookPrice(bid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"price": price})
}
