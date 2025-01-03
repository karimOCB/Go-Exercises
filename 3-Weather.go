// Package weather provides tools about the weather.
package weather

// CurrentCondition represents the current weather in certain place.
var CurrentCondition string
// CurrentLocation represents an specific place.
var CurrentLocation string

// Forecast return a string with a Location and its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
