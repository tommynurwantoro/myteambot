package method

import (
	"github.com/bot/act-bl-bot/text"
	"github.com/bot/act-bl-bot/utility"
	"github.com/bot/act-bl-bot/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// GroupChat _
func GroupChat(update tgbotapi.Update, groupSessionKey string, groupState int) string {
	if groupState == utility.RedisState["init"] {
		args := update.Message.CommandArguments()
		switch update.Message.CommandWithAt() {
		case "start":
			return text.Start()
		case "help":
			return text.Help()
		case "halo":
			return text.Halo(update.Message.From.UserName)
		case "retro":
			return "Kalau mau retro DM aku aja ya, biar anonymous hasilnya. 😄"
		case "result_retro":
			if args == "" {
				return text.InvalidRetroMessage()
			}
			results := mysql.GetResultRetro(args)
			return "Ini hasil retro untuk tanggal " + args + "\n\n" + text.GenerateRetroResult(results)
		default:
			return text.InvalidCommand()
		}
	}

	return text.InvalidCommand()
}
