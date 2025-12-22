package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate)*float64(successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var successRatePerHour float64 = CalculateWorkingCarsPerHour(productionRate,successRate)
    var result int = int(successRatePerHour/float64(60))
	return result
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var groupsOfTen int = carsCount/10
    var theRest int = carsCount % 10
    var result int = groupsOfTen*95000+theRest*10000
	return uint(result)
}
