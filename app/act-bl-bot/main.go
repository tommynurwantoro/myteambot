package main

import (
	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/app/method"
)

func main() {
	method.Init()

	app.Bot.Start()
}
