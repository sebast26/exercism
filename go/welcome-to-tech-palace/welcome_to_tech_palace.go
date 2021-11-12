package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var sb strings.Builder
	for n := 0; n < numStarsPerLine; n++ {
		sb.WriteString("*")
	}
	return fmt.Sprintf("%s\n%s\n%s", sb.String(), welcomeMsg, sb.String())
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msgSpaces := strings.ReplaceAll(oldMsg, "*", "")
	return strings.Trim(msgSpaces, "\n ")
}
