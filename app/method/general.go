package method

import (
	"os"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app"
	"github.com/bot/myteambot/app/utility"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Start _
func (m *Method) Start(message *tb.Message) {
	app.Bot.Send(message.Chat, utility.Start())
}

// Help _
func (m *Method) Help(message *tb.Message) {
	app.Bot.Send(message.Chat, utility.Help(utility.GenerateAllCommands()))
}

// Halo _
func (m *Method) Halo(message *tb.Message) {
	app.Bot.Send(message.Chat, utility.Halo(message.Sender.Username))
}

// GreetingFromBot _
func (m *Method) GreetingFromBot(message *tb.Message) {
	app.Bot.Send(message.Chat, utility.GreetingFromBot())
}

// GreetNewJoinedUser _
func (m *Method) GreetNewJoinedUser(message *tb.Message) {
	app.Bot.Send(message.Chat, utility.GreetingNewJoinedUser(message.UserJoined.Username))
}

// SendCustomChat _
func (m *Method) SendCustomChat(message *tb.Message) {
	if os.Getenv("ADMIN_USERNAME") != message.Sender.Username {
		m.Bot.Send(message.Chat, utility.UserNotEligible())
		return
	}

	if message.Payload == "" {
		app.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	split := strings.Split(message.Payload, "#")

	if len(split) < 2 {
		app.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	chatID := split[0]
	response := split[1]
	if chatID != "" {
		intChatID, err := strconv.ParseInt(chatID, 10, 64)
		if err != nil {
			app.Bot.Send(message.Chat, utility.InvalidParameter())
			return
		}
		message.Chat.ID = intChatID
	}

	app.Bot.Send(message.Chat, response)
}
