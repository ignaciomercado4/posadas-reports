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

func (h *ReportHandler) GetReports(c *gin.Context) {
	var existingReports []models.Report

	result := h.DB.Find(&existingReports)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.HTML(http.StatusOK, "reports.tmpl", gin.H{
		"title":   "Reports",
		"reports": existingReports,
	})
}

func (h *ReportHandler) CreateReport(c *gin.Context) {

	var reportInput models.ReportInput

	if err := c.ShouldBindJSON(&reportInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "binding error",
			"details": err.Error(),
		})
		return
	}

	newReport := models.Report{
		Category:    reportInput.Category,
		Title:       reportInput.Title,
		Description: reportInput.Description,
		PositionX:   reportInput.PositionX,
		PositionY:   reportInput.PositionY,
	}

	result := h.DB.Create(&newReport)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to create report",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"newReport": newReport,
	})
}
