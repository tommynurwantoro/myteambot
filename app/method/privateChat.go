package method

import (
	"log"
	"strconv"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/app/text"
	"github.com/bot/act-bl-bot/app/utility"
	"github.com/bot/act-bl-bot/app/utility/mysql"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// PrivateCommands List all private commands
var PrivateCommands = []Command{
	{"start", "Tentang bot ini"},
	{"help", "Nampilin semua perintah yang ada"},
	{"halo", "Cuma buat nyapa aja"},
	{"retro", "Masuk ke sesi retrospective"},
	{"result_retro", "{dd-mm-yyyy} Dapetin hasil retrospective, jangan lupa kasih tanggalnya ya"},
	{"titip_review", "{url} Titip review PR"},
	{"antrian_review", "Nampilin semua antrian PR yang belum direview"},
	{"sudah_direview", "{urutan} Ngubah antrian review untuk yang sudah direview"},
}

// PrivateChat _
func PrivateChat(update tgbotapi.Update) string {
	if !mysql.IsUserEligible(update.Message.From.UserName) {
		return text.UserNotEligible()
	}

	userSessionKey := "bot_user_session:" + update.Message.From.UserName
	if app.Redis.Exists(userSessionKey).Val() == 0 {
		err := app.Redis.Set(userSessionKey, utility.RedisState["init"], 0).Err()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}

	userState, err := strconv.Atoi(app.Redis.Get(userSessionKey).Val())
	if err != nil {
		panic(err)
	}

	args := update.Message.CommandArguments()

	if userState == utility.RedisState["init"] {
		switch update.Message.Command() {
		case PrivateCommands[0].Name: //start
			return text.Start()
		case PrivateCommands[1].Name: //help
			return text.Help(GenerateAllCommands(PrivateCommands))
		case PrivateCommands[2].Name: //halo
			return text.Halo(update.Message.From.UserName)
		case PrivateCommands[3].Name: //retro
			return StartRetro(update, userSessionKey)
		case PrivateCommands[4].Name: //result_retro
			return ResultRetro(args)
		case PrivateCommands[5].Name: //titip_review
			return AddReview(args)
		case PrivateCommands[6].Name: //antrian_review
			return GetReviewQueue()
		case PrivateCommands[7].Name: //sudah_direview
			return UpdateDoneReview(args)
		case "add_user":
			return AddUser(update, args)
		default:
			return text.InvalidCommand()
		}
	} else if userState == utility.RedisState["retro"] {
		switch update.Message.Command() {
		case "help":
			return text.StartRetro()
		case "glad":
			return InsertRetroMessage(update, "glad", args)
		case "sad":
			return InsertRetroMessage(update, "sad", args)
		case "mad":
			return InsertRetroMessage(update, "mad", args)
		case "result_retro":
			return ResultRetro(args)
		case "end_retro":
			err := app.Redis.Set(userSessionKey, utility.RedisState["init"], 0).Err()
			if err != nil {
				panic(err)
			}
			return text.EndRetro()
		default:
			return text.PleaseEndRetro()
		}
	}

	return text.InvalidCommand()
}
