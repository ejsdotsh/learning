package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * (successRate / float64(100)))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	cph := float64(productionRate) * (successRate / float64(100))

    return int(cph) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var totalCost uint
    perOneCar := 10000
    perTenCar := 95000
	numTenCar := carsCount / 10
    numOneCar := carsCount % 10
    
    if carsCount >= 10 {  
        totalCost = uint(numTenCar * perTenCar) + uint(numOneCar * perOneCar)
    } else if carsCount < 10 {
    	totalCost = uint(numOneCar * perOneCar)
    }
    
    return totalCost
}
