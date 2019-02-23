package mysql

import (
	"log"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/entity"
)

// GetAllNeedReviews _
func GetAllNeedReviews() []entity.Review {
	rows, err := app.MysqlClient.Query("SELECT * FROM reviews WHERE is_done = false ORDER BY created_at")
	if err != nil {
		log.Fatal(err)
	}

	results := make([]entity.Review, 0)
	for rows.Next() {
		var result entity.Review
		if err := rows.Scan(&result.ID, &result.URL, &result.IsDone, &result.CreatedAt, &result.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	return results
}

// GetAllReviewed _
func GetAllReviewed() []entity.Review {
	rows, err := app.MysqlClient.Query("SELECT * FROM reviews WHERE is_done = true ORDER BY created_at")
	if err != nil {
		log.Fatal(err)
	}

	results := make([]entity.Review, 0)
	for rows.Next() {
		var result entity.Review
		if err := rows.Scan(&result.ID, &result.URL, &result.IsDone, &result.CreatedAt, &result.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	return results
}

// InsertReview _
func InsertReview(url string) {
	_, err := app.MysqlClient.Exec(
		"INSERT INTO reviews(url, is_done, created_at, updated_at) VALUES(?, false, now(), now())",
		url)
	if err != nil {
		panic(err)
	}
}

// UpdateToDone _
func UpdateToDone(sequence int) bool {
	reviews := GetAllNeedReviews()

	for i, review := range reviews {
		if i == sequence-1 {
			_, err := app.MysqlClient.Exec("UPDATE reviews SET is_done = true WHERE id = ?", review.ID)
			if err != nil {
				panic(err)
			}

			return true
		}
	}

	return false
}
