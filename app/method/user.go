package method

import (
	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// AddUser _
func AddUser(username, args string, chatID int64) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
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
