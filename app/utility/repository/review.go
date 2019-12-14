package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/bot/myteambot/app/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GetAllNeedReview _
func GetAllNeedReview(groupID int) []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_done = ? AND users != '' AND group_id = ?", false, groupID), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// GetAllNeedQA _
func GetAllNeedQA(groupID int) []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_done = ? AND users = '' AND group_id = ?", false, groupID), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// GetAllDone _
func GetAllDone(groupID int) []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_done = ? AND group_id = ?", true, groupID), qm.OrderBy("created_at")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// InsertReview _
func InsertReview(title, url, users string, groupID int) {
	var review models.Review

	review.URL = url
	review.IsDone = false
	review.Title = title
	review.Users = users
	review.GroupID = groupID

	err := review.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

func UpdateReview(ID uint, title, url, users string) {
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
func UpdateToDoneReview(sequence int, groupID int, user string, force bool) bool {
	reviews := GetAllNeedReview(groupID)

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

func UpdateToReadyQA(sequences []string, groupID int64) bool {
	reviews := GetAllNeedReview(groupID)
	successToUpdate := false

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		for i, review := range reviews {
			if i+1 == sequence {
				review.IsReviewed = true

				err := review.UpdateG(boil.Infer())
				if err != nil {
					panic(err)
				}

				successToUpdate = true
				break
			}
		}
	}

	return successToUpdate
}

func UpdateToDoneQA(sequence int, groupID int) bool {
	reviews := GetAllNeedQA(groupID)

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

func removeAvailableUsers(users, deleteUser string) string {
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
