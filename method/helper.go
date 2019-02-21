package method

import (
	"regexp"
	"strings"
)

// GetUsernames _
func GetUsernames(usernames string) []string {
	arr := strings.Split(usernames, " ")
	newArr := make([]string, len(arr))
	reg, _ := regexp.Compile("[^0-9A-Za-z_]+")
	for i, username := range arr {
		username := reg.ReplaceAllString(username, "")
		newArr[i] = username
	}

	return newArr
}
