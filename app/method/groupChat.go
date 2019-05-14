package method

import (
	"strconv"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/app/text"
	"github.com/bot/act-bl-bot/app/utility"
	"github.com/bot/act-bl-bot/app/utility/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GroupCommands List all commands
var GroupCommands = []Command{
	{"start", "Tentang bot ini"},
	{"help", "Nampilin semua perintah yang ada"},
	{"halo", "Cuma buat nyapa aja"},
	{"retro", "Masuk ke sesi retrospective"},
	{"result_retro", "{dd-mm-yyyy} Dapetin hasil retrospective, jangan lupa kasih tanggalnya ya"},
	{"titip_review", "{title#url#telegram-users} Titip review PR"},
	{"antrian_review", "Nampilin semua antrian PR yang belum direview"},
	{"sudah_direview", "{urutan} Ngubah antrian review untuk yang sudah direview"},
	{"sudah_direview_semua", "{urutan} Ngubah antrian review untuk yang sudah direview untuk semua user"},
}

// GroupChat _
func GroupChat(update tgbotapi.Update) string {
	if !mysql.IsUserEligible(update.Message.From.UserName) {
		return text.UserNotEligible()
	}

	groupSessionKey := "bot_group_session:" + strconv.FormatInt(update.Message.Chat.ID, 10)

	// Set redis if key not exist
	if app.Redis.Exists(groupSessionKey).Val() == 0 {
		err := app.Redis.Set(groupSessionKey, utility.RedisState["init"], 0).Err()
		if err != nil {
			panic(err)
		}
	}

	groupState, err := strconv.Atoi(app.Redis.Get(groupSessionKey).Val())
	if err != nil {
		panic(err)
	}

	if groupState == utility.RedisState["init"] {
		args := update.Message.CommandArguments()
		switch update.Message.Command() {
		case GroupCommands[0].Name: //start
			return text.Start()
		case GroupCommands[1].Name: //help
			return text.Help(GenerateAllCommands(GroupCommands))
		case GroupCommands[2].Name: //halo
			return text.Halo(update.Message.From.UserName)
		case GroupCommands[3].Name: //retro
			return "Kalau mau retro DM aku aja ya, biar gak diliat yang lain. 😄"
		case GroupCommands[4].Name: //result_retro
			return ResultRetro(args)
		case GroupCommands[5].Name: //titip_review
			return AddReview(args)
		case GroupCommands[6].Name: //antrian_review
			return GetReviewQueue()
		case GroupCommands[7].Name: //sudah_direview
			return UpdateDoneReview(args, update.Message.From.UserName, false)
		case GroupCommands[8].Name: //sudah_direview_semua
			return UpdateDoneReview(args, update.Message.From.UserName, true)
		default:
			return text.InvalidCommand()
		}
	}

	return text.InvalidCommand()
}
