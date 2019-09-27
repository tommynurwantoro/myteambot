package method

import (
	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
	tb "gopkg.in/tucnak/telebot.v2"
)

type PrivateMethod struct {
	Bot     *tb.Bot
	Command *repository.Command
}

func (p *PrivateMethod) Glad(m *tb.Message) {
	p.Response(m, p.Command.Glad().Name)
}

func (p *PrivateMethod) Sad(m *tb.Message) {
	p.Response(m, p.Command.Sad().Name)
}

func (p *PrivateMethod) Mad(m *tb.Message) {
	p.Response(m, p.Command.Mad().Name)
}

// Response _
func (p *PrivateMethod) Response(m *tb.Message, command string) {
	if !m.Private() {
		p.Bot.Send(m.Chat, utility.RestrictGroupRetro())
		return
	}

	args := m.Payload
	c := p.Command

	switch command {
	case c.Glad().Name:
		p.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "glad", args))
	case c.Sad().Name:
		p.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "sad", args))
	case c.Mad().Name:
		p.Bot.Send(m.Sender, InsertRetroMessage(m.Sender.Username, "mad", args))
	}
}
