package method

import (
	"github.com/bot/act-bl-bot/app/text"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// AllCommands List all commands
var AllCommands = []Command{
	{"start", "Tentang bot ini"},
	{"help", "Nampilin semua perintah yang ada"},
	{"halo", "Cuma buat nyapa aja"},
	{"retro", "Masuk ke sesi retrospective"},
	{"result_retro", "{dd-mm-yyyy} Dapetin hasil retrospective, jangan lupa kasih tanggalnya ya"},
	{"titip_review", "{title#url#telegram-users} Titip review PR"},
	{"antrian_review", "Nampilin semua antrian PR yang belum direview"},
	{"sudah_direview", "{urutan} Ngubah antrian review untuk yang sudah direview"},
	{"sudah_direview_semua", "{urutan} Ngubah antrian review untuk yang sudah direview untuk semua user"},
	{"tambah_user_review", "{urutan#users} Nambahin user ke antrian review"},
}

// AllChat _
func AllChat(update tgbotapi.Update) string {
	args := update.Message.CommandArguments()
	switch update.Message.Command() {
	case AllCommands[0].Name: //start
		return text.Start()
	case AllCommands[1].Name: //help
		return text.Help(GenerateAllCommands(AllCommands))
	case AllCommands[2].Name: //halo
		return text.Halo(update.Message.From.UserName)
	case AllCommands[4].Name: //result_retro
		return ResultRetro(args)
	case AllCommands[5].Name: //titip_review
		return AddReview(args)
	case AllCommands[6].Name: //antrian_review
		return GetReviewQueue()
	case AllCommands[7].Name: //sudah_direview
		return UpdateDoneReview(args, update.Message.From.UserName, false)
	case AllCommands[8].Name: //sudah_direview_semua
		return UpdateDoneReview(args, update.Message.From.UserName, true)
	case AllCommands[9].Name: //tambah_user_review
		return AddUserReview(args)
	default:
		return ""
	}
}
