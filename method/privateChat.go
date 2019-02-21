package method

import (
	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/text"
	"github.com/bot/act-bl-bot/utility"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// PrivateChat _
func PrivateChat(update tgbotapi.Update, userSessionKey string, userState int) string {
	args := update.Message.CommandArguments()

	if userState == utility.RedisState["init"] {
		switch update.Message.Command() {
		case "start":
			return text.Start()
		case "help":
			return text.Help()
		case "halo":
			return text.Halo(update.Message.From.UserName)
		case "retro":
			return StartRetro(update, userSessionKey)
		case "result_retro":
			return ResultRetro(args)
		case "add_user":
			return AddUser(update, args)
		default:
			return text.InvalidCommand()
		}
	} else if userState == utility.RedisState["retro"] {
		switch update.Message.Command() {
		case "help":
			return text.StartRetro()
		case "glad":
			return InsertRetroMessage(update, "glad", args)
		case "sad":
			return InsertRetroMessage(update, "sad", args)
		case "mad":
			return InsertRetroMessage(update, "mad", args)
		case "result_retro":
			return ResultRetro(args)
		case "end_retro":
			err := app.Redis.Set(userSessionKey, utility.RedisState["init"], 0).Err()
			if err != nil {
				panic(err)
			}
			return text.EndRetro()
		default:
			return text.PleaseEndRetro()
		}
	}

	return text.InvalidCommand()
}
