package mysql

import (
	"database/sql"
	"log"
	"strings"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/bot/myteambot/app/models"
)

// GetAllNeedReview _
func GetAllNeedReview() []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_done = ? AND users != ''", false), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// GetAllNeedQA _
func GetAllNeedQA() []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_done = ? AND users = ''", false), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// GetAllDone _
func GetAllDone() []*models.Review {
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

func UpdateReview(ID uint, title string, url string, users string) {
	review, err := models.Reviews(qm.Where("id = ?", ID)).OneG()
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	review.URL = url
	review.IsDone = false
	review.Title = title
	review.Users = users

	err = review.UpdateG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// UpdateToDoneReview _
func UpdateToDoneReview(sequence int, user string, force bool) bool {
	reviews := GetAllNeedReview()

	for i, review := range reviews {
		if i == sequence-1 {
			if force {
				review.Users = ""
			} else {
				review.Users = removeAvailableUsers(review.Users, user)
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

func UpdateToDoneQA(sequence int) bool {
	reviews := GetAllNeedQA()

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
