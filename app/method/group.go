package method

import (
	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// AddGroup _
func AddGroup(chatID int, name string) string {
	group := mysql.FindGroupByChatID(chatID)
	if group == nil {
		mysql.UpsertGroup(chatID, name)
	}

	return text.SuccessInitGroup(name)
}
