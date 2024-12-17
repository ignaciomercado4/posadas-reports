package utils

import (
	"better-posadas/models"
	"math"
	"strings"
	"time"
)

func CalculateCategoryAmounts(reports []models.Report) map[string]int {
	categoryAmounts := map[string]int{
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
			categoryAmounts[category]++
		} else {
			categoryAmounts["other"]++
		}
	}

	return categoryAmounts
}

func CalculateUrgencyAmounts(reports []models.Report) map[string]int {
	urgencyAmounts := map[string]int{
		"none":     0,
		"low":      0,
		"medium":   0,
		"high":     0,
		"critical": 0,
	}

	for _, report := range reports {
		if report.Urgency == "None" {
			urgencyAmounts["none"]++
		}
		if report.Urgency == "Low" {
			urgencyAmounts["low"]++
		}
		if report.Urgency == "Medium" {
			urgencyAmounts["medium"]++
		}
		if report.Urgency == "High" {
			urgencyAmounts["high"]++
		}
		if report.Urgency == "Critical" {
			urgencyAmounts["critical"]++
		}
	}

	return urgencyAmounts
}

func CalculateRecentReportStats(reports []models.Report) map[string]int {
	recentStats := map[string]int{
		"last24Hours": 0,
		"last7Days":   0,
		"last30Days":  0,
	}

	now := time.Now()

	for _, report := range reports {
		daysSinceReport := now.Sub(report.CreatedAt).Hours() / 24
		if daysSinceReport <= 1 {
			recentStats["last24Hours"]++
		}
		if daysSinceReport <= 7 {
			recentStats["last7Days"]++
		}
		if daysSinceReport <= 30 {
			recentStats["last30Days"]++
		}
	}

	return recentStats
}

func CalculateCategoryPercentages(categoryAmounts map[string]int, totalReports int) map[string]float64 {
	percentages := make(map[string]float64)

	for category, count := range categoryAmounts {
		percentage := (float64(count) / float64(totalReports)) * 100
		percentages[category] = math.Round(percentage*100) / 100
	}

	return percentages
}
