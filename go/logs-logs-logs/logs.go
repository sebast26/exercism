package logs

import (
	"regexp"
	"strings"
	"fmt"
)

var r = regexp.MustCompile("\\[(?P<Level>[A-Z]+)\\]\\:\\W+(?P<Message>[\\w\\s\\W]+)")

// Message extracts the message from the provided log line.
func Message(line string) string {
	matches := r.FindStringSubmatch(line)
	return strings.TrimSpace(matches[2])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	message := Message(line)
	return len([]rune(message))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	matches := r.FindStringSubmatch(line)
	return strings.ToLower(matches[1])
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
