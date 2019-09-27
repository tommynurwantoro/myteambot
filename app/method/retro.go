package method

import (
	"fmt"
	"time"

	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
)

// InsertRetroMessage _
func InsertRetroMessage(username string, _type string, args string) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	if args == "" {
		return utility.InvalidRetroMessage()
	}

	user := repository.FindUserByUsername(username)

	repository.InsertMessageRetro(username, _type, args, user.GroupID)
	return utility.SuccessInsertMessage()
}

// GetResultRetro _
func GetResultRetro(username, args string) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	if args == "" {
		year, month, day := time.Now().Date()
		args = fmt.Sprintf("%d-%02d-%d", day, int(month), year)
	}

	user := repository.FindUserByUsername(username)

	results := repository.GetResultRetro(args, user.GroupID)
	return "Ini hasil retro untuk tanggal " + args + "\n\n" + utility.GenerateRetroResult(results)
}
