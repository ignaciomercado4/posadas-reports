package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "alive",
		})
	})

	r.Run(":" + string(PORT))
}
