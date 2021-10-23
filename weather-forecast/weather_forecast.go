/*
Package weather is a package that does something with the weather.
*/
package weather

//CurrentCondition shouldnt be in camelcase.
var CurrentCondition string

//CurrentLocation is another var.
var CurrentLocation string

//Forecast receives city and a condition, both strings, and returns a composed string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
