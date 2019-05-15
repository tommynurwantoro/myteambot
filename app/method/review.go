package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/act-bl-bot/app/text"
	"github.com/bot/act-bl-bot/app/utility/mysql"
)

// GetReviewQueue _
func GetReviewQueue() string {
	reviews := mysql.GetAllNeedReviews()

	if len(reviews) == 0 {
		return "Gak ada antrian review nih 👍🏻"
	}

	return fmt.Sprintf("Ini antrian review tim kamu:\n%s", GenerateAllNeedReviews(reviews))
}

// AddReview _
func AddReview(args string) string {
	if args == "" {
		return text.InvalidParameter()
	}

	split := strings.Split(args, "#")

	if len(split) < 3 {
		return text.InvalidParameter()
	}

	mysql.InsertReview(split[0], split[1], split[2])

	return text.SuccessInsertData()
}

// UpdateDoneReview _
func UpdateDoneReview(args string, username string, force bool) string {
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

		if mysql.UpdateToDone(sequence-updated, fmt.Sprintf("@%s", username), force) {
			updated++
			success = true
		}
	}

	if success {
		return fmt.Sprintf("%s\n%s", text.SuccessUpdateData(), GetReviewQueue())
	}

	return text.InvalidSequece()
}

func AddUserReview(args string) string {
	if args == "" {
		return text.InvalidParameter()
	}

	split := strings.Split(args, "#")

	sequence, err := strconv.Atoi(split[0])

	if len(split) < 2 || err != nil {
		return text.InvalidParameter()
	}

	reviews := mysql.GetAllNeedReviews()

	for i, review := range reviews {
		if i+1 == sequence {
			mysql.UpdateReview(review.ID, review.Title, review.URL, fmt.Sprintf("%s %s", review.Users, split[1]))
			return fmt.Sprintf("%s\n%s", text.SuccessUpdateData(), GetReviewQueue())
		}
	}

	return text.InvalidSequece()
}
