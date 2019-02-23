package method

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/bot/act-bl-bot/entity"
)

// Command _
type Command struct {
	Name        string
	Description string
}

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
func GenerateAllNeedReviews(reviews []entity.Review) string {
	var buffer bytes.Buffer

	for i, review := range reviews {
		buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, review.URL))
	}

	return buffer.String()
}

// GenerateAllCommands _
func GenerateAllCommands(commands []Command) string {
	var buffer bytes.Buffer

	for _, command := range commands {
		buffer.WriteString(fmt.Sprintf("/%s %s\n", command.Name, command.Description))
	}

	return buffer.String()
}
