package main

import (
	"better-posadas/models"
	"net/http"
	"strings"

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
		"title":   "Reports in Posadas",
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

func (h *ReportHandler) GetReportStats(c *gin.Context) {
	var existingReports []models.Report

	result := h.DB.Find(&existingReports)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	stats := gin.H{
		"categoryAmounts": gin.H{
			"infrastructure": 0,
			"environment":    0,
			"public safety":  0,
			"transportation": 0,
			"other":          0,
		},
	}

	for _, report := range existingReports {
		category := strings.ToLower(report.Category)

		if _, exists := stats["categoryAmounts"].(gin.H)[category]; exists {
			stats["categoryAmounts"].(gin.H)[category] =
				stats["categoryAmounts"].(gin.H)[category].(int) + 1
		} else {
			stats["categoryAmounts"].(gin.H)["other"] =
				stats["categoryAmounts"].(gin.H)["other"].(int) + 1
		}
	}

	c.HTML(http.StatusOK, "stats.tmpl", gin.H{
		"title": "Report statistics in Posadas",
		"stats": stats,
	})
}