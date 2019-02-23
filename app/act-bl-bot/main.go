package main

import (
	"log"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/method"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := app.Bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			log.Printf("%+v", update.Message.Chat)

			if update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup() {
				log.Printf("Group chat")
				msg.Text = method.GroupChat(update)
			} else {
				msg.Text = method.PrivateChat(update)
			}

			app.Bot.Send(msg)
		}
	}

	defer app.MysqlClient.Close()
}
