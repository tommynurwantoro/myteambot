package method

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/bot/myteambot/app/models"
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

// GenerateAllNeedReviews _
func GenerateAllNeedReviews(reviews []*models.Review) string {
	var buffer bytes.Buffer

	for i, review := range reviews {
		if review.Title == "" {
			buffer.WriteString(fmt.Sprintf("%d. <a href='%s'>Belum ada title</a> %s\n", i+1, review.URL, review.Users))
		} else {
			buffer.WriteString(fmt.Sprintf("%d. <a href='%s'>%s</a> %s\n", i+1, review.URL, review.Title, review.Users))
		}
	}

	return buffer.String()
}

// GenerateAllCommands _
func GenerateAllCommands(commands []Command) string {
	var buffer bytes.Buffer

	for _, command := range commands {
		buffer.WriteString(fmt.Sprintf("%s %s\n", command.Name, command.Description))
	}

	return buffer.String()
}
