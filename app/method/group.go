package method

import (
	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// AddGroup _
func AddGroup(chatID int64, name string) string {
	mysql.UpsertGroup(chatID, name)

	return text.SuccessInitGroup(name)
}
