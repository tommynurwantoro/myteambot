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
		"Bot ini punya beberapa fitur yang bisa kamu pakai untuk tim kamu.\n" +
		"Coba gunakan command /help untuk melihat perintah-perintah yang tersedia ya.\n" +
		"Want to contribute? https://github.com/tommynurwantoro/myteambot"
}

// Help _
func Help(commands string) string {
	return "Kamu bisa gunakan perintah-perintah ini loh:\n" + commands
}

// Halo _
func Halo(username string) string {
	return "Halo, @" + username + ". 👋🏻"
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

// CommandGroupOnly _
func CommandGroupOnly() string {
	return "Perintah ini cuma bisa di group ya"
}

// GreetingFromBot _
func GreetingFromBot() string {
	return "Halo semua, mohon bantuannya..."
}

// GreetingNewJoinedUser _
func GreetingNewJoinedUser(username string) string {
	return fmt.Sprintf("Welcome @%s!!! GLHF 😁", username)
}

// CustomCommandNotFound _
func CustomCommandNotFound() string {
	return "Belum ada custom command nih, pakai command /simpan_command dulu aja"
}
