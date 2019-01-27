package method

import (
	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/text"
	"github.com/bot/act-bl-bot/utility"
	"github.com/bot/act-bl-bot/utility/mysql"
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
			eligible := mysql.IsUserEligible(update.Message.From.UserName)
			if eligible == true {
				err := app.Redis.Set(userSessionKey, utility.RedisState["retro"], 0).Err()
				if err != nil {
					panic(err)
				}
				return text.StartRetro()
			}
			return "Kamu belum bisa ikutan retro, coba hubungin @tommynurwantoro"
		case "result_retro":
			if args == "" {
				return text.InvalidDate()
			}
			results := mysql.GetResultRetro(args)
			return "Ini hasil retro untuk tanggal " + args + "\n\n" + text.GenerateRetroResult(results)
		case "add_user":
			if args != "" {
				if mysql.IsAdmin(update.Message.From.UserName) {
					mysql.InsertOneUser(args)
					return "User " + args + " udah aku masukin nih biar bisa ikut retrospective juga kayak kamu."
				}

				return "Kamu gak boleh pakai perintah ini, ngomong dulu ke @tommynurwantoro ya"
			}
		default:
			return text.InvalidCommand()
		}
	} else if userState == utility.RedisState["retro"] {
		switch update.Message.Command() {
		case "help":
			return text.StartRetro()
		case "glad":
			if args == "" {
				return text.InvalidRetroMessage()
			}
			mysql.InsertMessageRetro(update.Message.From.UserName, "glad", args)
			return text.SuccessInsertMessage()
		case "sad":
			if args == "" {
				return text.InvalidRetroMessage()
			}
			mysql.InsertMessageRetro(update.Message.From.UserName, "sad", args)
			return text.SuccessInsertMessage()
		case "mad":
			if args == "" {
				return text.InvalidRetroMessage()
			}
			mysql.InsertMessageRetro(update.Message.From.UserName, "mad", args)
			return text.SuccessInsertMessage()
		case "result_retro":
			if args == "" {
				return text.InvalidDate()
			}
			results := mysql.GetResultRetro(args)
			return "Ini hasil retro untuk tanggal " + args + "\n\n" + text.GenerateRetroResult(results)
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
