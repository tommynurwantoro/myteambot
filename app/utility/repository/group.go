package repository

import (
	"database/sql"

	"github.com/bot/myteambot/app/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func InsertGroup(chatID int64, name string) {
	var group models.Group

	group.ChatID = chatID
	group.Name = name

	group.InsertG(boil.Infer())
}

func FindGroup(ID uint) *models.Group {
	group, err := models.Groups(qm.Where("id = ?", ID)).OneG()
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return group
}

func GetAllGroups() []*models.Group {
	groups, err := models.Groups().AllG()
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return groups
}

func FindGroupByChatID(chatID int64) *models.Group {
	group, err := models.Groups(qm.Where("chat_id = ?", chatID)).OneG()
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return group
}

func UpdateGroup(ID uint, name string) {
	group := FindGroup(ID)

	group.Name = name

	group.UpdateG(boil.Infer())
}

func UpsertGroup(chatID int64, name string) {
	group := FindGroupByChatID(chatID)
	if group == nil {
		InsertGroup(chatID, name)
	} else {
		UpdateGroup(group.ID, name)
	}
}
