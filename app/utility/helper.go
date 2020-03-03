package utility

import (
	"bytes"
	"fmt"

	"github.com/bot/myteambot/app/models"
	"github.com/bot/myteambot/app/utility/repository"
)

// GenerateHTMLReview _
func GenerateHTMLReview(reviews []*models.Review) string {
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
func GenerateAllCommands() string {
	var buffer bytes.Buffer

	for _, command := range repository.GetCommand().All() {
		buffer.WriteString(fmt.Sprintf("%s %s\n", command.Name, command.Description))
	}

	return buffer.String()
}

// GenerateCustomCommands _
func GenerateCustomCommands(commands []*models.CustomCommand) string {
	var buffer bytes.Buffer

	for i, command := range commands {
		buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, command.Command))
	}

	return buffer.String()
}
