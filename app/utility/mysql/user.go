package mysql

import (
	"log"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/bot/myteambot/app/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// FindUserByUsername _
func FindUserByUsername(username string) *models.User {
	user, err := models.Users(qm.Where("username = ?", username)).OneG()
	if err != nil {
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

// InsertOneUser _
func InsertOneUser(username string) {
	var user models.User

	user.Username = username

	err := user.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// FirstOrCreateUser _
func FirstOrCreateUser(username string) {
	user := FindUserByUsername(username)
	if user == nil {
		InsertOneUser(username)
	}
}
