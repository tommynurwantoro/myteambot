package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
)

func SaveCustomCommandGroup(chatID int64, username, args string) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	split := strings.Split(args, "#")

	if len(split) < 2 {
		return utility.InvalidParameter()
	}

	group := repository.FindGroupByChatID(chatID)

	repository.InsertCustomCommand(int(group.ID), split[0], split[1])

	return utility.SuccessInsertData()
}

func ListCustomCommandGroup(chatID int64, username string) string {
	if validation := IsValidRequest(username, "OK"); validation != "" {
		return validation
	}

	group := repository.FindGroupByChatID(chatID)
	customCommands := repository.GetAllCustomCommandsByGroupID(int(group.ID))

	if len(customCommands) == 0 {
		return utility.CustomCommandNotFound()
	}

	return fmt.Sprintf("Ini list command tim kamu:\n%s", utility.GenerateCustomCommands(customCommands))
}

func UpdateCustomCommandGroup(chatID int64, username, args string) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	split := strings.Split(args, "#")

	if len(split) < 2 {
		return utility.InvalidParameter()
	}

	sequence, err := strconv.Atoi(split[0])
	if err != nil {
		return utility.InvalidParameter()
	}

	group := repository.FindGroupByChatID(chatID)

	repository.UpdateCustomCommand(int(group.ID), sequence, split[1])

	return utility.SuccessUpdateData()
}

func DeleteCustomCommandGroup(chatID int64, username, args string) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	sequence, err := strconv.Atoi(args)
	if err != nil {
		return utility.InvalidParameter()
	}

	group := repository.FindGroupByChatID(chatID)

	repository.DeleteCustomCommand(int(group.ID), sequence)

	return utility.SuccessUpdateData()
}

func RespondCustomCommandGroup(chatID int64, args string) string {
	group := repository.FindGroupByChatID(chatID)
	if group == nil {
		return ""
	}

	commands := repository.GetAllCustomCommandsByGroupID(int(group.ID))
	if commands != nil {
		for _, c := range commands {
			if strings.Contains(args, c.Command) {
				return c.Message
			}
		}
	}

	return ""
}
