package method

import (
	"log"
	"strconv"

	"github.com/bot/myteambot/app"
	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
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
	app.Bot.Handle(c.InitGroup().Name, InitGroup)
	app.Bot.Handle(c.AddUser().Name, AddEligibleUser)
	app.Bot.Handle(tb.OnAddedToGroup, GreetingFromBot)
	app.Bot.Handle(tb.OnUserJoined, GreetNewJoinedUser)
	app.Bot.Handle(c.SendChat().Name, SendCustomChat)
	app.Bot.Handle(c.SimpanCommand().Name, SaveCommand)
	app.Bot.Handle(c.ListCommand().Name, ListCommand)
	app.Bot.Handle(c.UbahCommand().Name, UpdateCommand)
	app.Bot.Handle(c.HapusCommand().Name, DeleteCommand)
	app.Bot.Handle(c.BlastMessage().Name, BlastMessageToAllGroup)
	app.Bot.Handle(tb.OnText, RespondAllText)
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
	app.Bot.Send(m.Chat, GetResultRetro(m.Sender.Username, m.Payload))
}

func TitipReview(m *tb.Message) {
	app.Bot.Send(m.Chat, AddReview(m.Sender.Username, m.Payload))
}

func AntrianReview(m *tb.Message) {
	app.Bot.Send(m.Chat, GetReviewQueue(m.Sender.Username), tb.ModeHTML, tb.NoPreview)
}

func SudahDireview(m *tb.Message) {
	app.Bot.Send(m.Chat, UpdateDoneReview(m.Payload, m.Sender.Username, false), tb.ModeHTML, tb.NoPreview)
}

func SudahDireviewSemua(m *tb.Message) {
	app.Bot.Send(m.Chat, UpdateDoneReview(m.Payload, m.Sender.Username, true), tb.ModeHTML, tb.NoPreview)
}

func TambahUserReview(m *tb.Message) {
	app.Bot.Send(m.Chat, AddUserReview(m.Payload, m.Sender.Username), tb.ModeHTML, tb.NoPreview)
}

func AntrianQA(m *tb.Message) {
	app.Bot.Send(m.Chat, GetQAQueue(m.Sender.Username), tb.ModeHTML, tb.NoPreview)
}

func SudahDites(m *tb.Message) {
	app.Bot.Send(m.Chat, UpdateDoneQA(m.Payload, m.Sender.Username), tb.ModeHTML, tb.NoPreview)
}

func InitGroup(m *tb.Message) {
	if m.Private() {
		app.Bot.Send(m.Chat, text.CommandGroupOnly())
	} else {
		log.Println(m.Chat.ID)
		app.Bot.Send(m.Chat, AddGroup(m.Chat.ID, m.Chat.Title))
	}
}

func AddEligibleUser(m *tb.Message) {
	if m.Private() {
		app.Bot.Send(m.Chat, text.CommandGroupOnly())
	} else {
		app.Bot.Send(m.Chat, AddUser(m.Sender.Username, m.Payload, m.Chat.ID))
	}
}

func GreetingFromBot(m *tb.Message) {
	app.Bot.Send(m.Chat, text.GreetingFromBot())
}

func GreetNewJoinedUser(m *tb.Message) {
	app.Bot.Send(m.Chat, text.GreetingNewJoinedUser(m.UserJoined.Username))
}

func SendCustomChat(m *tb.Message) {
	chatID, response := SendChatToSpecificGroup(m.Sender.Username, m.Payload)
	if chatID != "" {
		intChatID, err := strconv.ParseInt(chatID, 10, 64)
		if err != nil {
			app.Bot.Send(m.Chat, text.InvalidParameter())
			return
		}
		m.Chat.ID = intChatID
	}
	app.Bot.Send(m.Chat, response)
}

func SaveCommand(m *tb.Message) {
	app.Bot.Send(m.Chat, SaveCustomCommandGroup(m.Chat.ID, m.Sender.Username, m.Payload))
}

func ListCommand(m *tb.Message) {
	app.Bot.Send(m.Chat, ListCustomCommandGroup(m.Chat.ID, m.Sender.Username))
}

func UpdateCommand(m *tb.Message) {
	app.Bot.Send(m.Chat, UpdateCustomCommandGroup(m.Chat.ID, m.Sender.Username, m.Payload))
}

func DeleteCommand(m *tb.Message) {
	app.Bot.Send(m.Chat, DeleteCustomCommandGroup(m.Chat.ID, m.Sender.Username, m.Payload))
}

func RespondAllText(m *tb.Message) {
	respond := RespondCustomCommandGroup(m.Chat.ID, m.Text)
	if respond != "" {
		app.Bot.Send(m.Chat, respond)
	}
}

func BlastMessageToAllGroup(m *tb.Message) {
	if !mysql.IsAdmin(m.Sender.Username) {
		app.Bot.Send(m.Chat, text.UserNotEligible())
		return
	}

	if m.Payload == "" {
		app.Bot.Send(m.Chat, text.InvalidParameter())
		return
	}

	allGroups := mysql.GetAllGroups()
	for _, group := range allGroups {
		m.Chat.ID = group.ChatID
		app.Bot.Send(m.Chat, m.Payload)
	}
}
