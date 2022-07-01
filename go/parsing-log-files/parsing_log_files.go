package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	validLine       = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	splitLine       = regexp.MustCompile(`<[~=*-]*>`)
	quotedPasswords = regexp.MustCompile(`"(?i)[\w\s]*password[\w\s]*"`)
	endOfLine       = regexp.MustCompile(`end-of-line\d+`)
	tagWithUserName = regexp.MustCompile(`User\s+(\w+)`)
)

func IsValidLine(text string) bool {
	return validLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitLine.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	sum := 0
	for _, line := range lines {
		if quotedPasswords.MatchString(line) {
			sum++
		}
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	return endOfLine.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	out := make([]string, len(lines))
	for i, line := range lines {
		if m := tagWithUserName.FindStringSubmatch(line); m != nil {
			out[i] = fmt.Sprintf("[USR] %s %s", m[1], line)
			continue
		}
		out[i] = line
	}
	return out
}
