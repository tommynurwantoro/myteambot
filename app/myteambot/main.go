package main

import (
	"github.com/bot/myteambot/app"
	"github.com/bot/myteambot/app/method"
)

func main() {
	method := method.NewMethod(app.Bot)
	method.InitCommand()

	app.Bot.Start()
}
