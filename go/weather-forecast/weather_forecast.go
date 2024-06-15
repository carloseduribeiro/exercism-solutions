// Package weather provides a weather forecast of various cities in the Goblinocus.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current location.
var CurrentLocation string

// Forecast returns the weather forecast for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
