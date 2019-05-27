package method

import (
	"github.com/bot/myteambot/app"
	"github.com/bot/myteambot/app/text"
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
	if !m.Private() {
		app.Bot.Send(m.Chat, text.RestrictGroupRetro())
		return
	}

	args := m.Payload
	c := GetCommand()

	switch command {
	case c.Glad().Name:
		app.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "glad", args))
	case c.Sad().Name:
		app.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "sad", args))
	case c.Mad().Name:
		app.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "mad", args))
	}
}
