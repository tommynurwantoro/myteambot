package method

import (
	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
)

func IsValidRequest(username, args string) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	if args == "" {
		return utility.InvalidParameter()
	}

	return ""
}

func IsValidGroup(chatID int64) bool {
	group := repository.FindGroupByChatID(chatID)
	if group == nil {
		return false
	}

	return true
}
