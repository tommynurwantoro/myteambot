package text

import (
	"strings"

	"github.com/bot/myteambot/app/models"
)

// UserNotEligible _
func UserNotEligible() string {
	return "Kamu belum bisa ikutan retro, coba hubungin @tommynurwantoro"
}

// Start _
func Start() string {
	return "Halo!!\nIni botnya tim ACT.\nCoba gunakan /help untuk melihat perintah-perintah yang tersedia ya.\nKalau ada yang bingung atau butuh akses coba hubungi @tommynurwantoro aja"
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

// StartRetro _
func StartRetro() string {
	return "Untuk sesi retrospective, silakan gunakan perintah di bawah ini yaa:\n" +
		"/glad pesan kamu\n" +
		"/sad pesan kamu\n" +
		"/mad pesan kamu\n\n" +
		"Tenang aja hasilnya anonymous kok.\n" +
		"Untuk mendapatkan hasilnya, kamu bisa gunakan perintah /result_retro dd-mm-yyyy\n"
}

// SuccessInsertMessage _
func SuccessInsertMessage() string {
	return "Pesan kamu udah aku catat ke database yaa.\nKalau mau aku catatin pesan lain juga boleh pake perintah yang sama kayak sebelumnya."
}

// RestrictGroupRetro _
func RestrictGroupRetro() string {
	return "Kalau mau gunakan perintah ini DM aku aja ya, biar gak diliat yang lain. 😄"
}

// InvalidRetroMessage _
func InvalidRetroMessage() string {
	return "Pesannya belum ada tuh. Coba lagi ya.."
}

// GenerateRetroResult _
func GenerateRetroResult(results []*models.Retro) string {
	glad := "Glad:\n"
	sad := "\nSad:\n"
	mad := "\nMad:\n"
	for _, result := range results {
		if result.Type == "mad" {
			mad += "- " + result.Message.String + "\n"
		} else if result.Type == "sad" {
			sad += "- " + result.Message.String + "\n"
		} else {
			glad += "- " + result.Message.String + "\n"
		}
	}

	return glad + sad + mad
}

// SuccessAddMember _
func SuccessAddMember(usernames []string) string {
	return "Berhasil menambahkan " + strings.Join(usernames[:], ", ")
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
