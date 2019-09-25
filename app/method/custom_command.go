package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

func SaveCustomCommandGroup(groupID int64, username, args string) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	split := strings.Split(args, "#")

	if len(split) < 2 {
		return text.InvalidParameter()
	}

	mysql.InsertCustomCommand(groupID, split[0], split[1])

	return text.SuccessInsertData()
}

func ListCustomCommandGroup(groupID int64, username string) string {
	if validation := IsValidRequest(username, "OK"); validation != "" {
		return validation
	}

	customCommands := mysql.GetAllCustomCommandsByGroupID(groupID)

	if len(customCommands) == 0 {
		return "Belum ada custom command nih, pakai command /simpan_command aja"
	}

	return fmt.Sprintf("Ini list command tim kamu:\n%s", GenerateCustomCommands(customCommands))
}

func UpdateCustomCommandGroup(groupID int64, username, args string) string {
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

	mysql.UpdateCustomCommand(groupID, sequence, split[1])

	return text.SuccessUpdateData()
}

func DeleteCustomCommandGroup(groupID int64, username, args string) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	sequence, err := strconv.Atoi(args)
	if err != nil {
		return text.InvalidParameter()
	}

	mysql.DeleteCustomCommand(groupID, sequence)

	return text.SuccessUpdateData()
}

func RespondCustomCommandGroup(groupID int64, args string) string {
	commands := mysql.GetAllCustomCommandsByGroupID(groupID)

	for _, c := range commands {
		if strings.Contains(args, c.Command) {
			return c.Message
		}
	}

	return ""
}
