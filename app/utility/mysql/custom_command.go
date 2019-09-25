package mysql

import (
	"database/sql"
	"log"

	"github.com/bot/myteambot/app/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func FindCutomCommand(commandID int64) *models.CustomCommand {
	command, err := models.CustomCommands(qm.Where("id = ?", commandID)).OneG()
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return command
}

func GetAllCustomCommandsByGroupID(groupID int64) []*models.CustomCommand {
	commands, err := models.CustomCommands(qm.Where("group_id = ?", groupID), qm.OrderBy("created_at")).AllG()
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return commands
}

func InsertCustomCommand(groupID int64, com, message string) {
	var command models.CustomCommand
	command.GroupID = groupID
	command.Command = com
	command.Message = message

	err := command.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

func UpdateCustomCommand(groupID int64, sequence int, message string) {
	commands := GetAllCustomCommandsByGroupID(int64(groupID))

	for i, c := range commands {
		if i == sequence-1 {
			c.Message = message
			c.UpdateG(boil.Infer())
		}
	}
}

func DeleteCustomCommand(groupID int64, sequence int) {
	commands := GetAllCustomCommandsByGroupID(int64(groupID))

	for i, c := range commands {
		if i == sequence-1 {
			c.DeleteG()
		}
	}
}
