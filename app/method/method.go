package method

import (
	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/app/text"
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
}

func Start(m *tb.Message) {
	app.Bot.Send(m.Sender, text.Start())
}

func Help(m *tb.Message) {
	app.Bot.Send(m.Sender, text.Help(GenerateAllCommands(GetCommand().All())))
}

func Halo(m *tb.Message) {
	app.Bot.Send(m.Sender, text.Halo(m.Sender.Username))
}

func Retro(m *tb.Message) {
	app.Bot.Send(m.Sender, text.StartRetro())
}

func ResultRetro(m *tb.Message) {
	app.Bot.Send(m.Sender, GetResultRetro(m.Payload))
}

func TitipReview(m *tb.Message) {
	app.Bot.Send(m.Sender, AddReview(m.Payload))
}

func AntrianReview(m *tb.Message) {
	app.Bot.Send(m.Sender, GetReviewQueue(), tb.ModeHTML)
}

func SudahDireview(m *tb.Message) {
	app.Bot.Send(m.Sender, UpdateDoneReview(m.Payload, m.Sender.Username, false), tb.ModeHTML)
}

func SudahDireviewSemua(m *tb.Message) {
	app.Bot.Send(m.Sender, UpdateDoneReview(m.Payload, m.Sender.Username, true), tb.ModeHTML)
}

func TambahUserReview(m *tb.Message) {
	app.Bot.Send(m.Sender, AddUserReview(m.Payload))
}
