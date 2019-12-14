package utility

import (
	"fmt"
)

// UserNotEligible _
func UserNotEligible() string {
	return "Kamu belum ada akses bot, coba hubungin @tommynurwantoro"
}

// Start _
func Start() string {
	return "Halo!!\n" +
		"Pertama-tama undang aku ke group kamu, terus gunakan perintah /add_user {user telegram teman-teman di group kamu}\n" +
		"Contoh: /add_user @haha @hehe @hihi\n" +
		"Udah deh kalian bisa pakai perintah-perintah yang lain.\n" +
		"Coba gunakan /help untuk melihat perintah-perintah yang tersedia ya.\n" +
		"Note : kamu cuma bisa pakai fitur di satu group ya, kalau udah punya group sebelumnya sama aku, ulangi perintah /add_user dengan parameter user kamu lagi.\n" +
		"\nKalau ada yang bingung atau butuh akses coba hubungi @tommynurwantoro aja"
}

// Help _
func Help(commands string) string {
	return "Kamu bisa gunakan perintah-perintah ini loh:\n" + commands
}

// Halo _
func Halo(username string) string {
	return "Halo, @" + username + ". 👋🏻"
}

// InvalidCommand _
func InvalidCommand() string {
	return "Aku gak ngerti perintah itu, coba perintah yang lain ya."
}

// InvalidDate _
func InvalidDate() string {
	return "Tanggalnya tolong dicek lagi ya, udah sesuai format dd-mm-yyyy belum?"
}

// InvalidParameter _
func InvalidParameter() string {
	return "Parameternya belum bener tuh, coba dicek lagi ya"
}

// SuccessInsertData _
func SuccessInsertData() string {
	return "OK!"
}

// SuccessUpdateData _
func SuccessUpdateData() string {
	return "Updated!"
}

// InvalidSequece _
func InvalidSequece() string {
	return "Gak bisa, gak ada di list"
}

// CheckPrivateMessage _
func CheckPrivateMessage() string {
	return "Cek DM kamu yaa"
}

func CommandGroupOnly() string {
	return "Perintah ini cuma bisa di group ya"
}

func GroupNotFound() string {
	return fmt.Sprintf("Group gak ketemu, pakai perintah /init_group dulu")
}

func GreetingFromBot() string {
	return "Halo semua, mohon bantuannya..."
}

func GreetingNewJoinedUser(username string) string {
	return fmt.Sprintf("Welcome @%s!!! GLHF 😁", username)
}

func CustomCommandNotFound() string {
	return "Belum ada custom command nih, pakai command /simpan_command dulu aja"
}
