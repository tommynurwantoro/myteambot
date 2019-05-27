package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/text"
	"github.com/bot/myteambot/app/utility/mysql"
)

// GetReviewQueue _
func GetReviewQueue(username string) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	user := mysql.FindUserByUsername(username)

	reviews := mysql.GetAllNeedReview(user.GroupID)

	if len(reviews) == 0 {
		return "Gak ada antrian review nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian review tim kamu:\n%s", GenerateHTMLReview(reviews))
}

// GetQAQueue _
func GetQAQueue(username string) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	user := mysql.FindUserByUsername(username)

	reviews := mysql.GetAllNeedQA(user.GroupID)

	if len(reviews) == 0 {
		return "Gak ada antrian QA nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian QA tim kamu:\n%s", GenerateHTMLReview(reviews))
}

// AddReview _
func AddReview(username, args string) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	if args == "" {
		return text.InvalidParameter()
	}

	user := mysql.FindUserByUsername(username)

	split := strings.Split(args, "#")

	if len(split) < 3 {
		return text.InvalidParameter()
	}

	mysql.InsertReview(split[0], split[1], split[2], user.GroupID)

	return text.SuccessInsertData()
}

// UpdateDoneReview _
func UpdateDoneReview(args, username string, force bool) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	user := mysql.FindUserByUsername(username)

	if args == "" {
		return text.InvalidParameter()
	}

	sequences := strings.Split(args, " ")
	success := false
	updated := 0

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		if mysql.UpdateToDoneReview(sequence-updated, user.GroupID, fmt.Sprintf("@%s", username), force) {
			updated++
			success = true
		}
	}

	if success {
		return fmt.Sprintf("%s\n%s", text.SuccessUpdateData(), GetReviewQueue(username))
	}

	return text.InvalidSequece()
}

// UpdateDoneQA _
func UpdateDoneQA(args, username string) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	user := mysql.FindUserByUsername(username)

	if args == "" {
		return text.InvalidParameter()
	}

	sequences := strings.Split(args, " ")
	success := false
	updated := 0

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		if mysql.UpdateToDoneQA(sequence-updated, user.GroupID) {
			updated++
			success = true
		}
	}

	if success {
		return fmt.Sprintf("%s\n%s", text.SuccessUpdateData(), GetQAQueue(username))
	}

	return text.InvalidSequece()
}

func AddUserReview(args, username string) string {
	if !mysql.IsUserEligible(username) {
		return text.UserNotEligible()
	}

	user := mysql.FindUserByUsername(username)

	if args == "" {
		return text.InvalidParameter()
	}

	split := strings.Split(args, "#")

	sequence, err := strconv.Atoi(split[0])

	if len(split) < 2 || err != nil {
		return text.InvalidParameter()
	}

	reviews := mysql.GetAllNeedReview(user.GroupID)

	for i, review := range reviews {
		if i+1 == sequence {
			mysql.UpdateReview(review.ID, review.Title, review.URL, fmt.Sprintf("%s %s", review.Users, split[1]))
			return fmt.Sprintf("%s\n%s", text.SuccessUpdateData(), GetReviewQueue(username))
		}
	}

	return text.InvalidSequece()
}
