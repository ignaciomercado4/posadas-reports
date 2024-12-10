package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReportHandler struct {
	DB *gorm.DB
}

func (h *ReportHandler) CreateReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Create report working",
	})
}
