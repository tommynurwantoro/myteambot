package mysql

import (
	"log"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/bot/act-bl-bot/app/models"
)

// GetAllNeedReviews _
func GetAllNeedReviews() []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_done = ?", false), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// GetAllReviewed _
func GetAllReviewed() []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_done = ?", true), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// InsertReview _
func InsertReview(url string) {
	var review models.Review

	review.URL = url
	review.IsDone = false

	err := review.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// UpdateToDone _
func UpdateToDone(sequence int) bool {
	reviews := GetAllNeedReviews()

	for i, review := range reviews {
		if i == sequence-1 {
			review.IsDone = true
			err := review.UpdateG(boil.Infer())
			if err != nil {
				panic(err)
			}

			return true
		}
	}

	return false
}
