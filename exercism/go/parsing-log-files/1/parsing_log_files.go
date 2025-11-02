package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	if len(re.FindString(text)) > 0 {
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`\".*(?i)password.*\"`)
	cnt := 0
	for _, line := range lines {
		if len(re.FindString(line)) > 0 {
			cnt++
		}
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w{6,}) `)
	tagd := []string{}
	for _, line := range lines {
		if len(re.FindString(line)) == 0 {
			tagd = append(tagd, line)
		} else {
			tagdLog := "[USR] " + strings.Fields(re.FindString(line))[1] + " " + line
			tagd = append(tagd, tagdLog)
		}
	}
	return tagd
}
