package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
)

// GetReviewQueue _
func GetReviewQueue(groupID int64) string {
	reviews := repository.GetAllNeedReview(groupID)

	if len(reviews) == 0 {
		return "Gak ada antrian review nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian review tim kamu:\n%s", utility.GenerateHTMLReview(reviews))
}

// GetQAQueue _
func GetQAQueue(groupID int64) string {
	reviews := repository.GetAllNeedQA(groupID)

	if len(reviews) == 0 {
		return "Gak ada antrian QA nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian QA tim kamu:\n%s", utility.GenerateHTMLReview(reviews))
}

// AddReview _
func AddReview(groupID int64, args string) string {
	if args == "" {
		return utility.InvalidParameter()
	}

	split := strings.Split(args, "#")

	if len(split) < 3 {
		return utility.InvalidParameter()
	}

	repository.InsertReview(split[0], split[1], split[2], groupID)

	return utility.SuccessInsertData()
}

// UpdateDoneReview _
func UpdateDoneReview(groupID int64, username, args string, force bool) string {
	if args == "" {
		return utility.InvalidParameter()
	}

	sequences := strings.Split(args, " ")
	success := repository.UpdateToDoneReview(sequences, groupID, fmt.Sprintf("@%s", username), force)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetReviewQueue(groupID))
	}

	return utility.InvalidSequece()
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
func UpdateDoneQA(groupID int64, args string) string {
	if args == "" {
		return utility.InvalidParameter()
	}

	sequences := strings.Split(args, " ")
	success := repository.UpdateToDoneQA(sequences, groupID)

	if success {
		return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetQAQueue(groupID))
	}

	return utility.InvalidSequece()
}

func AddUserReview(groupID int64, args string) string {
	if args == "" {
		return utility.InvalidParameter()
	}

	split := strings.Split(args, "#")

	sequence, err := strconv.Atoi(split[0])

	if len(split) < 2 || err != nil {
		return utility.InvalidParameter()
	}

	reviews := repository.GetAllNeedReview(groupID)

	for i, review := range reviews {
		if i+1 == sequence {
			repository.UpdateReview(review.ID, review.Title, review.URL, fmt.Sprintf("%s %s", review.Users, split[1]))
			return fmt.Sprintf("%s\n%s", utility.SuccessUpdateData(), GetReviewQueue(groupID))
		}
	}

	return utility.InvalidSequece()
}
