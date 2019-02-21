package method

import (
	"github.com/bot/act-bl-bot/text"
	"github.com/bot/act-bl-bot/utility"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// GroupChat _
func GroupChat(update tgbotapi.Update, groupSessionKey string, groupState int) string {
	if groupState == utility.RedisState["init"] {
		args := update.Message.CommandArguments()
		switch update.Message.Command() {
		case "start":
			return text.Start()
		case "help":
			return text.Help()
		case "halo":
			return text.Halo(update.Message.From.UserName)
		case "retro":
			return "Kalau mau retro DM aku aja ya, biar gak diliat yang lain. 😄"
		case "result_retro":
			return ResultRetro(args)
		default:
			return text.InvalidCommand()
		}
	}

	return text.InvalidCommand()
}
