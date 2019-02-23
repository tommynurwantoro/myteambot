package method

import (
	"strconv"

	"github.com/bot/act-bl-bot/text"
	"github.com/bot/act-bl-bot/utility/mysql"
)

// GetReviewQueue _
func GetReviewQueue() string {
	reviews := mysql.GetAllNeedReviews()

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
	sequence, err := strconv.Atoi(args)
	if err != nil {
		return text.InvalidParameter()
	}

	success := mysql.UpdateToDone(sequence)

	if success {
		return text.SuccessUpdateData()
	}

	return text.InvalidSequece()
}
