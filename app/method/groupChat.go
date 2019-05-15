package method

import (
	"strconv"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/app/text"
	"github.com/bot/act-bl-bot/app/utility"
	"github.com/bot/act-bl-bot/app/utility/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GroupChat _
func GroupChat(update tgbotapi.Update) string {
	if !mysql.IsUserEligible(update.Message.From.UserName) {
		return text.UserNotEligible()
	}

	groupSessionKey := "bot_group_session:" + strconv.FormatInt(update.Message.Chat.ID, 10)

	// Set redis if key not exist
	if app.Redis.Exists(groupSessionKey).Val() == 0 {
		err := app.Redis.Set(groupSessionKey, utility.RedisState["init"], 0).Err()
		if err != nil {
			panic(err)
		}
	}

	groupState, err := strconv.Atoi(app.Redis.Get(groupSessionKey).Val())
	if err != nil {
		panic(err)
	}

	if groupState == utility.RedisState["init"] {
		switch update.Message.Command() {
		case AllCommands[3].Name: //retro
			return "Kalau mau retro DM aku aja ya, biar gak diliat yang lain. 😄"
		default:
			return text.InvalidCommand()
		}
	}

	return text.InvalidCommand()
}
