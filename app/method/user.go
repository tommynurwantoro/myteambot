package method

import (
	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// AddUser _
func AddUser(username string, args string) string {
	if args == "" {
		return text.InvalidParameter()
	}

	if !mysql.IsAdmin(username) {
		return "Kamu gak boleh pakai perintah ini, ngomong dulu ke @tommynurwantoro ya"
	}

	usernames := GetUsernames(args)
	for _, username := range usernames {
		mysql.FirstOrCreateUser(username)
	}

	return text.SuccessAddMember(usernames)
}
