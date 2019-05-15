package method

import (
	"log"
	"strconv"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/app/text"
	"github.com/bot/act-bl-bot/app/utility"
	"github.com/bot/act-bl-bot/app/utility/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// PrivateChat _
func PrivateChat(update tgbotapi.Update) string {
	if !mysql.IsUserEligible(update.Message.From.UserName) {
		return text.UserNotEligible()
	}

	userSessionKey := "bot_user_session:" + update.Message.From.UserName
	if app.Redis.Exists(userSessionKey).Val() == 0 {
		err := app.Redis.Set(userSessionKey, utility.RedisState["init"], 0).Err()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}

	userState, err := strconv.Atoi(app.Redis.Get(userSessionKey).Val())
	if err != nil {
		panic(err)
	}

	args := update.Message.CommandArguments()

	if userState == utility.RedisState["init"] {
		switch update.Message.Command() {
		case AllCommands[3].Name: //retro
			return StartRetro(update, userSessionKey)
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
