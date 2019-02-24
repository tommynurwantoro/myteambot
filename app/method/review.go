package method

import (
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

	return "Ini antrian review tim kamu:\n" + GenerateAllNeedReviews(reviews)
}

// AddReview _
func AddReview(url string) string {
	if url == "" {
		return text.InvalidParameter()
	}

	mysql.InsertReview(url)

	return text.SuccessInsertData()
}

// UpdateDoneReview _
func UpdateDoneReview(args string) string {
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

		if mysql.UpdateToDone(sequence - updated) {
			updated++
			success = true
		}
	}

	if success {
		return text.SuccessUpdateData()
	}

	return text.InvalidSequece()
}
