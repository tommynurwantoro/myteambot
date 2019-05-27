package method

import (
	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// AddUser _
func AddUser(username, args string, chatID int64) string {
	if args == "" {
		return text.InvalidParameter()
	}

	if !mysql.IsAdmin(username) {
		return "Kamu gak boleh pakai perintah ini, ngomong dulu ke @tommynurwantoro ya"
	}

	group := mysql.FindGroupByChatID(chatID)
	if group == nil {
		return text.GroupNotFound()
	}

	usernames := GetUsernames(args)
	for _, username := range usernames {
		mysql.UpsertUser(username, int(group.ID))
	}

	return text.SuccessAddMember(usernames)
}
