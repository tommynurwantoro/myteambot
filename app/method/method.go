package method

import (
	"strconv"

	"github.com/bot/myteambot/app"
	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
	tb "gopkg.in/tucnak/telebot.v2"
)

func Init() {
	var command *repository.Command

	app.Bot.Handle(command.Start().Name, Start)
	app.Bot.Handle(command.Help().Name, Help)
	app.Bot.Handle(command.Halo().Name, Halo)
	app.Bot.Handle(command.TitipReview().Name, TitipReview)
	app.Bot.Handle(command.AntrianReview().Name, AntrianReview)
	app.Bot.Handle(command.SudahDireview().Name, SudahDireview)
	app.Bot.Handle(command.SudahDireviewSemua().Name, SudahDireviewSemua)
	app.Bot.Handle(command.TambahUserReview().Name, TambahUserReview)
	app.Bot.Handle(command.SiapQA().Name, SiapQA)
	app.Bot.Handle(command.AntrianQA().Name, AntrianQA)
	app.Bot.Handle(command.SudahDites().Name, SudahDites)
	app.Bot.Handle(command.SendChat().Name, SendCustomChat)
	app.Bot.Handle(command.SimpanCommand().Name, SaveCommand)
	app.Bot.Handle(command.ListCommand().Name, ListCommand)
	app.Bot.Handle(command.UbahCommand().Name, UpdateCommand)
	app.Bot.Handle(command.HapusCommand().Name, DeleteCommand)
	app.Bot.Handle(command.BlastMessage().Name, BlastMessageToAllGroup)
	app.Bot.Handle(tb.OnAddedToGroup, GreetingFromBot)
	app.Bot.Handle(tb.OnUserJoined, GreetNewJoinedUser)
	app.Bot.Handle(tb.OnText, RespondAllText)
}

func Start(m *tb.Message) {
	app.Bot.Send(m.Chat, utility.Start())
}

func Help(m *tb.Message) {
	app.Bot.Send(m.Chat, utility.Help(utility.GenerateAllCommands()))
}

func Halo(m *tb.Message) {
	app.Bot.Send(m.Chat, utility.Halo(m.Sender.Username))
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

func SiapQA(m *tb.Message) {
	invalid := validateGroup(m)

	if invalid != "" {
		app.Bot.Send(m.Chat, invalid)
		return
	}

	app.Bot.Send(m.Chat, UpdateReadyQA(m.Chat.ID, m.Payload), tb.ModeHTML, tb.NoPreview)
}

func AntrianQA(m *tb.Message) {
	app.Bot.Send(m.Chat, GetQAQueue(m.Sender.Username), tb.ModeHTML, tb.NoPreview)
}

func SudahDites(m *tb.Message) {
	app.Bot.Send(m.Chat, UpdateDoneQA(m.Payload, m.Sender.Username), tb.ModeHTML, tb.NoPreview)
}

func GreetingFromBot(m *tb.Message) {
	AddGroup(m.Chat.ID, m.Chat.Title)
	app.Bot.Send(m.Chat, utility.GreetingFromBot())
}

func GreetNewJoinedUser(m *tb.Message) {
	app.Bot.Send(m.Chat, utility.GreetingNewJoinedUser(m.UserJoined.Username))
}

func SendCustomChat(m *tb.Message) {
	chatID, response := SendChatToSpecificGroup(m.Sender.Username, m.Payload)
	if chatID != "" {
		intChatID, err := strconv.ParseInt(chatID, 10, 64)
		if err != nil {
			app.Bot.Send(m.Chat, utility.InvalidParameter())
			return
		}
		m.Chat.ID = intChatID
	}
	app.Bot.Send(m.Chat, response)
}

func SaveCommand(m *tb.Message) {
	if m.Private() {
		app.Bot.Send(m.Chat, utility.CommandGroupOnly())
	} else if !IsValidGroup(m.Chat.ID) {
		app.Bot.Send(m.Chat, utility.GroupNotFound())
	} else {
		app.Bot.Send(m.Chat, SaveCustomCommandGroup(m.Chat.ID, m.Sender.Username, m.Payload))
	}
}

func ListCommand(m *tb.Message) {
	if m.Private() {
		app.Bot.Send(m.Chat, utility.CommandGroupOnly())
	} else if !IsValidGroup(m.Chat.ID) {
		app.Bot.Send(m.Chat, utility.GroupNotFound())
	} else {
		app.Bot.Send(m.Chat, ListCustomCommandGroup(m.Chat.ID, m.Sender.Username))
	}
}

func UpdateCommand(m *tb.Message) {
	if m.Private() {
		app.Bot.Send(m.Chat, utility.CommandGroupOnly())
	} else if !IsValidGroup(m.Chat.ID) {
		app.Bot.Send(m.Chat, utility.GroupNotFound())
	} else {
		app.Bot.Send(m.Chat, UpdateCustomCommandGroup(m.Chat.ID, m.Sender.Username, m.Payload))
	}
}

func DeleteCommand(m *tb.Message) {
	if m.Private() {
		app.Bot.Send(m.Chat, utility.CommandGroupOnly())
	} else if !IsValidGroup(m.Chat.ID) {
		app.Bot.Send(m.Chat, utility.GroupNotFound())
	} else {
		app.Bot.Send(m.Chat, DeleteCustomCommandGroup(m.Chat.ID, m.Sender.Username, m.Payload))
	}
}

func RespondAllText(m *tb.Message) {
	respond := RespondCustomCommandGroup(m.Chat.ID, m.Text)
	if respond != "" {
		app.Bot.Send(m.Chat, respond)
	}
}

func BlastMessageToAllGroup(m *tb.Message) {
	if !repository.IsAdmin(m.Sender.Username) {
		app.Bot.Send(m.Chat, utility.UserNotEligible())
		return
	}

	if m.Payload == "" {
		app.Bot.Send(m.Chat, utility.InvalidParameter())
		return
	}

	allGroups := repository.GetAllGroups()
	for _, group := range allGroups {
		m.Chat.ID = group.ChatID
		app.Bot.Send(m.Chat, m.Payload)
	}
}

func validateGroup(m *tb.Message) string {
	if m.Private() {
		return utility.CommandGroupOnly()
	} else if !IsValidGroup(m.Chat.ID) {
		return utility.GroupNotFound()
	}

	return ""
}
