package main

import (
	"better-posadas/models"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

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
	existingReports, err := h.fetchReports()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stats := h.compileReportStatistics(existingReports)

	c.HTML(http.StatusOK, "stats.tmpl", gin.H{
		"title": "Comprehensive Report Statistics in Posadas",
		"stats": stats,
	})
}

func (h *ReportHandler) fetchReports() ([]models.Report, error) {
	var existingReports []models.Report
	result := h.DB.Find(&existingReports)
	if result.Error != nil {
		return nil, result.Error
	}
	return existingReports, nil
}

func (h *ReportHandler) compileReportStatistics(reports []models.Report) gin.H {
	stats := gin.H{
		"categoryAmounts":     h.calculateCategoryAmounts(reports),
		"totalReports":        len(reports),
		"recentStats":         h.calculateRecentReportStats(reports),
		"categoryPercentages": make(gin.H),
	}

	stats["categoryPercentages"] = h.calculateCategoryPercentages(
		stats["categoryAmounts"].(gin.H),
		stats["totalReports"].(int),
	)

	return stats
}

func (h *ReportHandler) calculateCategoryAmounts(reports []models.Report) gin.H {

	categoryAmounts := gin.H{
		"infrastructure":         0,
		"environment":            0,
		"public safety":          0,
		"transportation":         0,
		"garbage disposal":       0,
		"road marking/signaling": 0,
		"vandalism":              0,
		"other":                  0,
	}

	for _, report := range reports {
		category := strings.ToLower(report.Category)

		if _, exists := categoryAmounts[category]; exists {
			categoryAmounts[category] = categoryAmounts[category].(int) + 1
		} else {
			categoryAmounts["other"] = categoryAmounts["other"].(int) + 1
		}
	}

	return categoryAmounts
}

func (h *ReportHandler) calculateRecentReportStats(reports []models.Report) gin.H {
	recentStats := gin.H{
		"last24Hours": 0,
		"last7Days":   0,
		"last30Days":  0,
	}

	now := time.Now()

	fmt.Println("Recent Reports Debug:")
	for _, report := range reports {
		daysSinceReport := now.Sub(report.CreatedAt)
		fmt.Printf("Report CreatedAt: %v, Days Since: %f\n", report.CreatedAt, daysSinceReport.Hours()/24)

		if daysSinceReport.Hours() <= 24 {
			recentStats["last24Hours"] = recentStats["last24Hours"].(int) + 1
		}
		if daysSinceReport.Hours() <= 24*7 {
			recentStats["last7Days"] = recentStats["last7Days"].(int) + 1
		}
		if daysSinceReport.Hours() <= 24*30 {
			recentStats["last30Days"] = recentStats["last30Days"].(int) + 1
		}
	}

	return recentStats
}

func (h *ReportHandler) calculateCategoryPercentages(
	categoryAmounts gin.H,
	totalReports int,
) gin.H {
	percentages := make(gin.H)

	for category, count := range categoryAmounts {
		percentage := (float64(count.(int)) / float64(totalReports)) * 100
		percentages[category] = math.Round(percentage*100) / 100
	}

	return percentages
}
