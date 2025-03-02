package main

import (
	"log"
	"net/http"

	"github.com/junietoc/best-stocks-server/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.GET("/stocks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	log.Println("server is running on port 8080")
	r.Run(":8080")
}
