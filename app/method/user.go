package method

import (
	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
)

// AddUser _
func AddUser(username, args string, chatID int64) string {
	if validation := IsValidRequest(username, args); validation != "" {
		return validation
	}

	group := repository.FindGroupByChatID(chatID)
	if group == nil {
		return utility.GroupNotFound()
	}

	usernames := utility.GetUsernames(args)
	for _, username := range usernames {
		repository.UpsertUser(username, int(group.ID))
	}

	return utility.SuccessAddMember(usernames)
}
