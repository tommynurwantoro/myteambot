package method

import (
	"github.com/bot/act-bl-bot/app/text"
	"github.com/bot/act-bl-bot/app/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// AddUser _
func AddUser(update tgbotapi.Update, args string) string {
	if args == "" {
		return text.InvalidParameter()
	}

	if !mysql.IsAdmin(update.Message.From.UserName) {
		return "Kamu gak boleh pakai perintah ini, ngomong dulu ke @tommynurwantoro ya"
	}

	usernames := GetUsernames(args)
	for _, username := range usernames {
		mysql.FirstOrCreateUser(username)
	}

	return text.SuccessAddMember(usernames)
}
