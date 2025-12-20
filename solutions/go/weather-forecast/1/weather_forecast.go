// Package weather checks the current weather conditions
// in th curent location and outputs the results.
package weather

var (
    // CurrentCondition represets the curret condition.
	CurrentCondition string
    // CurrentLocation represets the curret location.
	CurrentLocation  string
)

// Forecast outputs the values of variables CurrentCondition and CurrentLocaton.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
