package mysql

import (
	"log"
	"strings"

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
func InsertReview(title string, url string, users string) {
	var review models.Review

	review.URL = url
	review.IsDone = false
	review.Title = title
	review.Users = users

	err := review.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// UpdateToDone _
func UpdateToDone(sequence int, user string, force bool) bool {
	reviews := GetAllNeedReviews()

	for i, review := range reviews {
		if i == sequence-1 {
			if force {
				review.Users = ""
			} else {
				review.Users = removeAvailableUsers(review.Users, user)
			}

			if review.Users == "" {
				review.IsDone = true
			}

			err := review.UpdateG(boil.Infer())
			if err != nil {
				panic(err)
			}

			return true
		}
	}

	return false
}

func removeAvailableUsers(users string, deleteUser string) string {
	splitUsers := strings.Split(users, " ")
	var newUsers []string

	for _, user := range splitUsers {
		if user == deleteUser {
			continue
		} else {
			newUsers = append(newUsers, user)
		}
	}

	return strings.Join(newUsers, " ")
}
