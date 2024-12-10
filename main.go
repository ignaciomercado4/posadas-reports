package main

import (
	"better-posadas/database"
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

	r := gin.Default()

	db := database.ConnectDatabase()
	database.MigrateModels(db)

	reportHandler := ReportHandler{DB: db}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "alive",
		})
	})

	r.GET("/createReport", reportHandler.CreateReport)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	r.Run(":" + string(PORT))
}
