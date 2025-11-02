// Package weather provides a function to return the forecast
// given a city and condition.
package weather

// CurrentCondition represents the current weather conditions.
var CurrentCondition string
// CurrentLocation represents the current location, or city, for the weather.
var CurrentLocation string

// Forecast returns the joined strings of the city and the condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
