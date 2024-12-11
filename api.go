package main

import (
	"better-posadas/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReportHandler struct {
	DB *gorm.DB
}

func (h *ReportHandler) CreateReport(c *gin.Context) {
	var reportInput models.ReportInput

	if err := c.ShouldBind(&reportInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "binding error in CreateReport",
		})
		return
	}

	newReport := models.Report{
		Category:    reportInput.Category,
		Title:       reportInput.Title,
		Description: reportInput.Category,
		PositionX:   reportInput.PositionX,
		PositionY:   reportInput.PositionY,
	}

	h.DB.Create(&newReport)
	c.JSON(http.StatusOK, gin.H{
		"newReport": newReport,
	})
}
