package utils

import "time"

func CalculateTotalDays(startRent, endRent time.Time) int {
	if startRent.Equal(endRent) {
		return 1
	}
	return int(endRent.Sub(startRent).Hours()/24) + 1
}

func CalculateDiscount(daysOfRent int, dailyRentPrice, discountRate float64) float64 {
	totalRent := float64(daysOfRent) * dailyRentPrice
	discount := totalRent * discountRate / 100
	return discount
}

func CalculateDriverCost(daysOfRent int, dailyCost float64) float64 {
	totalRent := float64(daysOfRent) * dailyCost
	return totalRent
}

func CalculateDriverIncentive(daysOfRent int, dailyRentPrice float64) float64 {
	totalRent := float64(daysOfRent) * dailyRentPrice
	incentive := totalRent * 0.05
	return incentive
}
