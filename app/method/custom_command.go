package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

func SaveCustomCommandGroup(chatID int64, username, args string) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	split := strings.Split(args, "#")

	if len(split) < 2 {
		return text.InvalidParameter()
	}

	group := mysql.FindGroupByChatID(chatID)

	mysql.InsertCustomCommand(int(group.ID), split[0], split[1])

	return text.SuccessInsertData()
}

func ListCustomCommandGroup(chatID int64, username string) string {
	if validation := IsValidRequest(username, "OK"); validation != "" {
		return validation
	}

	group := mysql.FindGroupByChatID(chatID)
	customCommands := mysql.GetAllCustomCommandsByGroupID(int(group.ID))

	if len(customCommands) == 0 {
		return "Belum ada custom command nih, pakai command /simpan_command aja"
	}

	return fmt.Sprintf("Ini list command tim kamu:\n%s", GenerateCustomCommands(customCommands))
}

func UpdateCustomCommandGroup(chatID int64, username, args string) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	split := strings.Split(args, "#")

	if len(split) < 2 {
		return text.InvalidParameter()
	}

	sequence, err := strconv.Atoi(split[0])
	if err != nil {
		return text.InvalidParameter()
	}

	group := mysql.FindGroupByChatID(chatID)

	mysql.UpdateCustomCommand(int(group.ID), sequence, split[1])

	return text.SuccessUpdateData()
}

func DeleteCustomCommandGroup(chatID int64, username, args string) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	sequence, err := strconv.Atoi(args)
	if err != nil {
		return text.InvalidParameter()
	}

	group := mysql.FindGroupByChatID(chatID)

	mysql.DeleteCustomCommand(int(group.ID), sequence)

	return text.SuccessUpdateData()
}

func RespondCustomCommandGroup(chatID int64, args string) string {
	group := mysql.FindGroupByChatID(chatID)
	commands := mysql.GetAllCustomCommandsByGroupID(int(group.ID))

	for _, c := range commands {
		if strings.Contains(args, c.Command) {
			return c.Message
		}
	}

	return ""
}
