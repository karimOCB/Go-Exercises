package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stringOfStars := strings.Repeat("*", numStarsPerLine)
    return stringOfStars + "\n" + welcomeMsg + "\n" + stringOfStars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	noStars := strings.ReplaceAll(oldMsg, "*", "")
    cleanedMsg := strings.TrimSpace(noStars)

    return cleanedMsg
}
