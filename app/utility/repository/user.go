package repository

import (
	"database/sql"
	"log"

	"github.com/bot/myteambot/app/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// FindUserByUsername _
func FindUserByUsername(username string) *models.User {
	user, err := models.Users(qm.Where("username = ?", username)).OneG()
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}

	return user
}

// IsUserEligible _
func IsUserEligible(username string) bool {
	user := FindUserByUsername(username)
	if user == nil {
		return false
	}

	return true
}

// IsAdmin _
func IsAdmin(username string) bool {
	user := FindUserByUsername(username)
	if user == nil {
		return false
	}

	return user.IsAdmin
}

// InsertUser _
func InsertUser(username string, groupID int) {
	var user models.User

	user.Username = username
	user.GroupID = groupID

	err := user.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// UpdateUser _
func UpdateUser(username string, groupID int) {
	user := FindUserByUsername(username)

	user.Username = username
	user.GroupID = groupID

	err := user.UpdateG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// UpsertUser _
func UpsertUser(username string, groupID int) {
	user := FindUserByUsername(username)
	if user == nil {
		InsertUser(username, groupID)
	} else {
		UpdateUser(username, groupID)
	}
}
