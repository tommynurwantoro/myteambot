package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
	tb "gopkg.in/tucnak/telebot.v2"
)

// SimpanCustomCommand _
func (m *Method) SimpanCustomCommand(message *tb.Message) {
	if invalid := ValidateGroup(message); invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter(), tb.ModeHTML, tb.NoPreview)
		return
	}

	split := strings.Split(message.Payload, "#")

	if len(split) < 2 {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	repository.InsertCustomCommand(message.Chat.ID, split[0], split[1])

	m.Bot.Send(message.Chat, utility.SuccessInsertData())
}

// ListCustomCommand _
func (m *Method) ListCustomCommand(message *tb.Message) {
	if invalid := ValidateGroup(message); invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	customCommands := repository.GetAllCustomCommandsByGroupID(message.Chat.ID)

	if len(customCommands) == 0 {
		m.Bot.Send(message.Chat, utility.CustomCommandNotFound())
		return
	}

	m.Bot.Send(message.Chat, fmt.Sprintf("Ini list command tim kamu:\n%s", utility.GenerateCustomCommands(customCommands)))
}

// UbahCustomCommand _
func (m *Method) UbahCustomCommand(message *tb.Message) {
	if invalid := ValidateGroup(message); invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter(), tb.ModeHTML, tb.NoPreview)
		return
	}

	split := strings.Split(message.Payload, "#")

	if len(split) < 2 {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	sequence, err := strconv.Atoi(split[0])
	if err != nil {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	repository.UpdateCustomCommand(message.Chat.ID, sequence, split[1])

	m.Bot.Send(message.Chat, utility.SuccessUpdateData())
}

// HapusCustomCommand _
func (m *Method) HapusCustomCommand(message *tb.Message) {
	if invalid := ValidateGroup(message); invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter(), tb.ModeHTML, tb.NoPreview)
		return
	}

	sequence, err := strconv.Atoi(message.Payload)
	if err != nil {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	repository.DeleteCustomCommand(message.Chat.ID, sequence)

	m.Bot.Send(message.Chat, utility.SuccessUpdateData())
}

// RespondCustomCommand _
func (m *Method) RespondCustomCommand(message *tb.Message) {
	commands := repository.GetAllCustomCommandsByGroupID(message.Chat.ID)
	if commands != nil {
		for _, c := range commands {
			if strings.Contains(strings.ToLower(message.Payload), strings.ToLower(c.Command)) {
				m.Bot.Send(message.Chat, c.Message)
				return
			}
		}
	}
}
