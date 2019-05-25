package method

import (
	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/app/text"
	"github.com/bot/act-bl-bot/app/utility/mysql"
	tb "gopkg.in/tucnak/telebot.v2"
)

type PrivateMethod struct {
	Bot *tb.Bot
}

func NewPrivateMethod() *PrivateMethod {
	return &PrivateMethod{
		Bot: app.Bot,
	}
}

func (p *PrivateMethod) AddUser(m *tb.Message) {
	p.Response(m, "add_user")
}

func (p *PrivateMethod) Glad(m *tb.Message) {
	p.Response(m, GetCommand().Glad().Name)
}

func (p *PrivateMethod) Sad(m *tb.Message) {
	p.Response(m, GetCommand().Sad().Name)
}

func (p *PrivateMethod) Mad(m *tb.Message) {
	p.Response(m, GetCommand().Mad().Name)
}

// Response _
func (p *PrivateMethod) Response(m *tb.Message, command string) {
	if !mysql.IsUserEligible(m.Sender.Username) {
		app.Bot.Send(m.Sender, text.UserNotEligible())
		return
	}

	if !m.Private() {
		app.Bot.Send(m.Sender, text.RestrictGroupRetro())
		return
	}

	args := m.Payload
	c := GetCommand()

	switch command {
	case "add_user":
		app.Bot.Send(m.Sender, AddUser(m.Sender.Username, args))
	case c.Glad().Name:
		app.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "glad", args))
	case c.Sad().Name:
		app.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "sad", args))
	case c.Mad().Name:
		app.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "mad", args))
	}
}
