package mysql

import (
	"log"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/entity"
)

// GetOneUser _
func GetOneUser(username string) entity.User {
	user := entity.User{}
	err := app.MysqlClient.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.IsAdmin)
	if err != nil {
		log.Println(err)
	}

	return user
}

// IsUserEligible _
func IsUserEligible(username string) bool {
	user := GetOneUser(username)
	if user == (entity.User{}) {
		return false
	}

	return true
}

// IsAdmin _
func IsAdmin(username string) bool {
	user := GetOneUser(username)
	if user == (entity.User{}) {
		return false
	}

	return user.IsAdmin
}

// InsertOneUser _
func InsertOneUser(username string) {
	_, err := app.MysqlClient.Exec(
		"INSERT INTO users(username, is_admin) VALUES(?, ?)",
		username, false)
	if err != nil {
		panic(err)
	}
}
