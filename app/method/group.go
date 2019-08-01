package method

import (
	"strings"

	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// AddGroup _
func AddGroup(chatID int64, name string) string {
	mysql.UpsertGroup(chatID, name)

	return text.SuccessInitGroup(name)
}

// SendChatSpecificGroup _
func SendChatToSpecificGroup(username, args string) (string, string) {
	if !mysql.IsUserEligible(username) {
		return "", text.UserNotEligible()
	}

	if args == "" {
		return "", text.InvalidParameter()
	}

	split := strings.Split(args, "#")

	if len(split) < 2 {
		return "", text.InvalidParameter()
	}

	return split[0], split[1]
}
