package handlers

import (
	"better-posadas/models"
	"better-posadas/utils"
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
		Urgency:     reportInput.Urgency,
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
		"title": "Report Statistics in Posadas",
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
		"categoryAmounts":     utils.CalculateCategoryAmounts(reports),
		"urgencyAmounts":      utils.CalculateUrgencyAmounts(reports),
		"totalReports":        len(reports),
		"recentStats":         utils.CalculateRecentReportStats(reports),
		"categoryPercentages": make(gin.H),
	}

	stats["categoryPercentages"] = utils.CalculateCategoryPercentages(
		stats["categoryAmounts"].(map[string]int),
		stats["totalReports"].(int),
	)

	return stats
}
