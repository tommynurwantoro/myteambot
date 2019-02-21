package text

import (
	"strings"

	"github.com/bot/act-bl-bot/entity"
)

// Start _
func Start() string {
	return "Halo!!\nCoba gunakan /help untuk melihat perintah-perintah yang tersedia.\nKalau ada yang bingung coba hubungi @tommynurwantoro aja"
}

// Help _
func Help() string {
	return "Kamu bisa gunakan perintah-perintah ini loh:\n" + commands()
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
	return "Kamu memasuki sesi retrospective nih.\n" +
		"Silakan gunakan perintah di bawah ini yaa:\n" +
		"/glad pesan kamu\n" +
		"/sad pesan kamu\n" +
		"/mad pesan kamu\n\n" +
		"Tenang aja hasilnya anonymous kok.\n" +
		"Untuk mendapatkan hasilnya, kamu bisa gunakan perintah /result_retro dd-mm-yyyy\n" +
		"Untuk menghentikan sesi retro ini, kamu bisa gunakan perintah /end_retro`\n"
}

// SuccessInsertMessage _
func SuccessInsertMessage() string {
	return "Pesan kamu udah aku catat ke database yaa.\nKalau mau aku catatin pesan lain juga boleh pake perintah yang sama kayak sebelumnya."
}

// InvalidRetroMessage _
func InvalidRetroMessage() string {
	return "Pesannya belum ada tuh. Coba lagi ya.."
}

// PleaseEndRetro _
func PleaseEndRetro() string {
	return "Selesaikan sesi retronya dulu ya dengan perintah /end_retro"
}

// EndRetro _
func EndRetro() string {
	return "Terimakasih, sesi retronya udah berakhir."
}

// GenerateRetroResult _
func GenerateRetroResult(results []entity.Retro) string {
	glad := "Glad:\n"
	sad := "\nSad:\n"
	mad := "\nMad:\n"
	for _, result := range results {
		if result.Type == "mad" {
			mad += "- " + result.Message + "\n"
		} else if result.Type == "sad" {
			sad += "- " + result.Message + "\n"
		} else {
			glad += "- " + result.Message + "\n"
		}
	}

	return glad + sad + mad
}

// SuccessAddMember _
func SuccessAddMember(usernames []string) string {
	return "Berhasil menambahkan " + strings.Join(usernames[:], ", ")
}

// Private //

func commands() string {
	return "/halo Cuma buat nyapa aja\n" +
		"/retro Buat masuk ke sesi retrospective\n" +
		"/result_retro dd-mm-yyyy Buat dapet hasil retrospective, jangan lupa kasih tanggalnya ya"
}
