package main

import (
	"github.com/bot/myteambot/app"
	"github.com/bot/myteambot/app/method"
)

func main() {
	method.Init()

	app.Bot.Start()
}
