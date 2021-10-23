package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
// 0: 0% success rate.
// 1 - 4: 100% success rate.
// 5 - 8: 90% success rate.
// 9 - 10: 77% success rate.
func SuccessRate(speed int) float64 {
	if speed == 0 {
		return 0
	} else if speed >= 1 && speed <= 4 {
		return 1
	} else if speed >= 5 && speed <= 8 {
		return 0.9
	} else if speed >= 9 && speed <= 10 {
		return 0.77
	} else {
		return 0
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return (SuccessRate(speed) * 221 * float64(speed))
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {

	return int(CalculateProductionRatePerHour(speed)) / 60
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	if CalculateProductionRatePerHour(speed) > limit {
		return limit
	} else {
		return CalculateProductionRatePerHour(speed)
	}
}
