package method

import (
	"fmt"
	"time"

	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// InsertRetroMessage _
func InsertRetroMessage(username string, _type string, args string) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	if args == "" {
		return text.InvalidRetroMessage()
	}

	user := mysql.FindUserByUsername(username)

	mysql.InsertMessageRetro(username, _type, args, user.GroupID)
	return text.SuccessInsertMessage()
}

// GetResultRetro _
func GetResultRetro(username, args string) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	if args == "" {
		year, month, day := time.Now().Date()
		args = fmt.Sprintf("%d-%02d-%d", day, int(month), year)
	}

	user := mysql.FindUserByUsername(username)

	results := mysql.GetResultRetro(args, user.GroupID)
	return "Ini hasil retro untuk tanggal " + args + "\n\n" + text.GenerateRetroResult(results)
}
