package method

import (
	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

func IsValidRequest(username, args string) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	if args == "" {
		return text.InvalidParameter()
	}

	return ""
}
