package main

import (
	"log"
	"net/http"

	"best-stocks-server/internal/db"
	"best-stocks-server/internal/stocks"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/fetch-and-save", func(c *gin.Context) {
		items, err := stocks.FetchStockData()
		if err != nil {
			log.Println("Error fetching stock data:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = stocks.SaveStockItems(db.DB, items)
		if err != nil {
			log.Println("Error saving stock data:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Stock data fetched and saved successfully!",
			"count":   len(items),
		})
	})

	r.GET("/stocks", func(c *gin.Context) {
		stocksData, err := db.GetStocks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stocksData)
	})

	log.Println("server is running on port 8080")
	r.Run(":8080")
}
