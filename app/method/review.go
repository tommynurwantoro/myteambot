package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
)

// GetReviewQueue _
func GetReviewQueue(username string) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	user := repository.FindUserByUsername(username)

	reviews := repository.GetAllNeedReview(user.GroupID)

	if len(reviews) == 0 {
		return "Gak ada antrian review nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian review tim kamu:\n%s", utility.GenerateHTMLReview(reviews))
}

// GetQAQueue _
func GetQAQueue(username string) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	user := repository.FindUserByUsername(username)

	reviews := repository.GetAllNeedQA(user.GroupID)

	if len(reviews) == 0 {
		return "Gak ada antrian QA nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian QA tim kamu:\n%s", utility.GenerateHTMLReview(reviews))
}

// AddReview _
func AddReview(username, args string) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	if args == "" {
		return utility.InvalidParameter()
	}

	user := repository.FindUserByUsername(username)

	split := strings.Split(args, "#")

	if len(split) < 3 {
		return utility.InvalidParameter()
	}

	repository.InsertReview(split[0], split[1], split[2], user.GroupID)

	return utility.SuccessInsertData()
}

// UpdateDoneReview _
func UpdateDoneReview(args, username string, force bool) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	user := repository.FindUserByUsername(username)

	if args == "" {
		return utility.InvalidParameter()
	}

	sequences := strings.Split(args, " ")
	success := false
	updated := 0

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		if repository.UpdateToDoneReview(sequence-updated, user.GroupID, fmt.Sprintf("@%s", username), force) {
			updated++
			success = true
		}

// UpdateReadyQA _
func UpdateReadyQA(groupID int64, args string) string {
	if args == "" {
		return utility.InvalidParameter()
	}

	sequences := strings.Split(args, " ")
	success := repository.UpdateToReadyQA(sequences, groupID)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetReviewQueue(groupID))
	}

	return utility.InvalidSequece()
}

// UpdateDoneQA _
func UpdateDoneQA(args, username string) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	user := repository.FindUserByUsername(username)

	if args == "" {
		return utility.InvalidParameter()
	}

	sequences := strings.Split(args, " ")
	success := false
	updated := 0

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		if repository.UpdateToDoneQA(sequence-updated, user.GroupID) {
			updated++
			success = true
		}
	}

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetQAQueue(username))
	}

	return utility.InvalidSequece()
}

func AddUserReview(args, username string) string {
	if !repository.IsUserEligible(username) {
		return utility.UserNotEligible()
	}

	user := repository.FindUserByUsername(username)

	if args == "" {
		return utility.InvalidParameter()
	}

	split := strings.Split(args, "#")

	sequence, err := strconv.Atoi(split[0])

	if len(split) < 2 || err != nil {
		return utility.InvalidParameter()
	}

	reviews := repository.GetAllNeedReview(user.GroupID)

	for i, review := range reviews {
		if i+1 == sequence {
			repository.UpdateReview(review.ID, review.Title, review.URL, fmt.Sprintf("%s %s", review.Users, split[1]))
			return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetReviewQueue(username))
		}
	}

	return utility.InvalidSequece()
}
