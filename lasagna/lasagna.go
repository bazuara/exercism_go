package lasagna

// OvenTime returns the expected baking time in minutes.
func OvenTime() int {
	return 40
}

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(time int) int {
	return OvenTime() - time
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(layer int) int {
	return 2 * layer
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(layers int, time int) int {
	return PreparationTime(layers) + time
}
