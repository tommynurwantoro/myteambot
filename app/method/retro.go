package method

import (
	"fmt"
	"time"

	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// InsertRetroMessage _
func InsertRetroMessage(username string, _type string, args string) string {
	if args == "" {
		return text.InvalidRetroMessage()
	}
	mysql.InsertMessageRetro(username, _type, args)
	return text.SuccessInsertMessage()
}

// GetResultRetro _
func GetResultRetro(args string) string {
	if args == "" {
		year, month, day := time.Now().Date()
		args = fmt.Sprintf("%d-%02d-%d", day, int(month), year)
	}

	results := mysql.GetResultRetro(args)
	return "Ini hasil retro untuk tanggal " + args + "\n\n" + text.GenerateRetroResult(results)
}
