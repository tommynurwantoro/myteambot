package method

import (
	"github.com/bot/myteambot/app/utility/repository"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Method _
type Method struct {
	Bot *tb.Bot
}

// NewMethod _
func NewMethod(bot *tb.Bot) *Method {
	return &Method{Bot: bot}
}

// InitCommand _
func (m *Method) InitCommand() {
	var command *repository.Command

	m.Bot.Handle(command.Start().Name, m.Start)
	m.Bot.Handle(command.Help().Name, m.Help)
	m.Bot.Handle(command.Halo().Name, m.Halo)
	m.Bot.Handle(command.TitipReview().Name, m.TitipReview)
	m.Bot.Handle(command.AntrianReview().Name, m.AntrianReview)
	m.Bot.Handle(command.SudahDireview().Name, m.SudahDireview)
	m.Bot.Handle(command.SudahDireviewSemua().Name, m.SudahDireviewSemua)
	m.Bot.Handle(command.TambahUserReview().Name, m.TambahUserReview)
	m.Bot.Handle(command.HapusReview().Name, m.HapusReview)
	m.Bot.Handle(command.SiapQA().Name, m.SiapQA)
	m.Bot.Handle(command.AntrianQA().Name, m.AntrianQA)
	m.Bot.Handle(command.SudahDites().Name, m.SudahDites)
	m.Bot.Handle(command.SendChat().Name, m.SendCustomChat)
	m.Bot.Handle(command.SimpanCommand().Name, m.SimpanCustomCommand)
	m.Bot.Handle(command.ListCommand().Name, m.ListCustomCommand)
	m.Bot.Handle(command.UbahCommand().Name, m.UbahCustomCommand)
	m.Bot.Handle(command.HapusCommand().Name, m.HapusCustomCommand)
	m.Bot.Handle(tb.OnAddedToGroup, m.GreetingFromBot)
	m.Bot.Handle(tb.OnUserJoined, m.GreetNewJoinedUser)
	m.Bot.Handle(tb.OnText, m.RespondCustomCommand)
}
