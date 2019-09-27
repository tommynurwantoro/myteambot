package method

import (
	"strings"

	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
)

// AddGroup _
func AddGroup(chatID int64, name string) string {
	repository.UpsertGroup(chatID, name)

	return utility.SuccessInitGroup(name)
}

// SendChatSpecificGroup _
func SendChatToSpecificGroup(username, args string) (string, string) {
	if !repository.IsUserEligible(username) {
		return "", utility.UserNotEligible()
	}

	if args == "" {
		return "", utility.InvalidParameter()
	}

	split := strings.Split(args, "#")

	if len(split) < 2 {
		return "", utility.InvalidParameter()
	}

	return split[0], split[1]
}
