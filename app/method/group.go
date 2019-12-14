package method

import (
	"strings"

	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
)

// AddGroup _
func AddGroup(chatID int64, name string) {
	repository.UpsertGroup(chatID, name)
}

// SendChatSpecificGroup _
func SendChatToSpecificGroup(args string) (string, string) {
	if args == "" {
		return "", utility.InvalidParameter()
	}

	split := strings.Split(args, "#")

	if len(split) < 2 {
		return "", utility.InvalidParameter()
	}

	return split[0], split[1]
}
