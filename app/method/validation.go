package method

import (
	"github.com/bot/myteambot/app/utility"
	tb "gopkg.in/tucnak/telebot.v2"
)

// ValidateGroup _
func ValidateGroup(message *tb.Message) string {
	if message.Private() {
		return utility.CommandGroupOnly()
	}

	return ""
}
