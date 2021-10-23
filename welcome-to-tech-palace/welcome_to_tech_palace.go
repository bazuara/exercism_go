package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
// Output: Welcome to the Tech Palace, JUDY
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
// Output:
// **********
// Welcome!
// **********
func AddBorder(welcomeMsg string, numStarsPerLine int) string {

	line := strings.Repeat("*", numStarsPerLine)
	return line + "\n" + welcomeMsg + "\n" + line
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.Split(oldMsg, "\n")
	line := newMsg[1]
	line = strings.TrimSuffix(line, "*")
	line = strings.TrimPrefix(line, "*")
	line = strings.TrimSpace(line)
	return line
}
