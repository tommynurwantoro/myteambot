package method

import (
	"github.com/bot/myteambot/app"
	"github.com/bot/myteambot/app/text"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	private *PrivateMethod
)

func Init() {
	c := GetCommand()
	private = NewPrivateMethod()

	app.Bot.Handle(c.Start().Name, Start)
	app.Bot.Handle(c.Help().Name, Help)
	app.Bot.Handle(c.Halo().Name, Halo)
	app.Bot.Handle(c.Retro().Name, Retro)
	app.Bot.Handle(c.Glad().Name, private.Glad)
	app.Bot.Handle(c.Sad().Name, private.Sad)
	app.Bot.Handle(c.Mad().Name, private.Mad)
	app.Bot.Handle(c.ResultRetro().Name, ResultRetro)
	app.Bot.Handle(c.TitipReview().Name, TitipReview)
	app.Bot.Handle(c.AntrianReview().Name, AntrianReview)
	app.Bot.Handle(c.SudahDireview().Name, SudahDireview)
	app.Bot.Handle(c.SudahDireviewSemua().Name, SudahDireviewSemua)
	app.Bot.Handle(c.TambahUserReview().Name, TambahUserReview)
	app.Bot.Handle(c.AntrianQA().Name, AntrianQA)
	app.Bot.Handle(c.SudahDites().Name, SudahDites)
}

func Start(m *tb.Message) {
	app.Bot.Send(m.Chat, text.Start())
}

func Help(m *tb.Message) {
	app.Bot.Send(m.Chat, text.Help(GenerateAllCommands(GetCommand().All())))
}

func Halo(m *tb.Message) {
	app.Bot.Send(m.Chat, text.Halo(m.Sender.Username))
}

func Retro(m *tb.Message) {
	if !m.Private() {
		app.Bot.Send(m.Chat, text.CheckPrivateMessage())
	}
	app.Bot.Send(m.Sender, text.StartRetro())
}

func ResultRetro(m *tb.Message) {
	app.Bot.Send(m.Chat, GetResultRetro(m.Payload))
}

func TitipReview(m *tb.Message) {
	app.Bot.Send(m.Chat, AddReview(m.Payload))
}

func AntrianReview(m *tb.Message) {
	app.Bot.Send(m.Chat, GetReviewQueue(), tb.ModeHTML)
}

func SudahDireview(m *tb.Message) {
	app.Bot.Send(m.Chat, UpdateDoneReview(m.Payload, m.Sender.Username, false), tb.ModeHTML)
}

func SudahDireviewSemua(m *tb.Message) {
	app.Bot.Send(m.Chat, UpdateDoneReview(m.Payload, m.Sender.Username, true), tb.ModeHTML)
}

func TambahUserReview(m *tb.Message) {
	app.Bot.Send(m.Chat, AddUserReview(m.Payload))
}

func AntrianQA(m *tb.Message) {
	app.Bot.Send(m.Chat, GetQAQueue(), tb.ModeHTML)
}

func SudahDites(m *tb.Message) {
	app.Bot.Send(m.Chat, UpdateDoneQA(m.Payload))
}
