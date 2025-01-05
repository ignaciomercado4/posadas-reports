package main

import (
	"better-posadas/database"
	"better-posadas/handlers"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using system environment variables")
		}
	}

	r := gin.Default()

	r.Use(cors.Default())

	db := database.ConnectDatabase()
	database.MigrateModels(db)

	reportHandler := handlers.ReportHandler{DB: db}

	r.SetFuncMap(template.FuncMap{
		"replace": strings.ReplaceAll,
		"add":     func(a, b int) int { return a + b },
	})

	r.StaticFile("/reportsUi.js", "./public/reportsUi.js")
	r.StaticFile("/mapSearchBar.js", "./public/mapSearchBar.js")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.LoadHTMLGlob("./templates/*") // fix

	r.GET("/", reportHandler.GetHome)
	r.GET("/reports", reportHandler.GetReports)
	r.GET("/reports/:id", reportHandler.GetReportDetail)
	r.POST("/reports/create", reportHandler.CreateReport)
	r.GET("/reports/stats", reportHandler.GetReportStats)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Server is running on port " + PORT)
	r.Run(":" + PORT)
}
