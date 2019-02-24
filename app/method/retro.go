package method

import (
	"fmt"
	"time"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/app/text"
	"github.com/bot/act-bl-bot/app/utility"
	"github.com/bot/act-bl-bot/app/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartRetro _
func StartRetro(update tgbotapi.Update, userSessionKey string) string {
	eligible := mysql.IsUserEligible(update.Message.From.UserName)
	if eligible == true {
		err := app.Redis.Set(userSessionKey, utility.RedisState["retro"], 0).Err()
		if err != nil {
			panic(err)
		}
		return text.StartRetro()
	}
	return text.UserNotEligible()
}

// InsertRetroMessage _
func InsertRetroMessage(update tgbotapi.Update, _type string, args string) string {
	if args == "" {
		return text.InvalidRetroMessage()
	}
	mysql.InsertMessageRetro(update.Message.From.UserName, _type, args)
	return text.SuccessInsertMessage()
}

// ResultRetro _
func ResultRetro(args string) string {
	if args == "" {
		year, month, day := time.Now().Date()
		args = fmt.Sprintf("%d-%02d-%d", day, int(month), year)
	}

	results := mysql.GetResultRetro(args)
	return "Ini hasil retro untuk tanggal " + args + "\n\n" + text.GenerateRetroResult(results)
}
