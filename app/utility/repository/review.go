package repository

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// GetAllNeedReview _
func GetAllNeedReview(groupID int64) []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_reviewed = ? AND group_id = ?", false, groupID), qm.OrderBy("created_at"), qm.Limit(30)).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// GetAllNeedQA _
func GetAllNeedQA(groupID int64) []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_reviewed = ? AND is_tested = ? AND group_id = ?", true, false, groupID), qm.OrderBy("created_at"), qm.Limit(30)).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// GetAllDone _
func GetAllDone(groupID int64) []*models.Review {
	reviews, err := models.Reviews(qm.Where("is_tested = ? AND group_id = ?", true, groupID), qm.OrderBy("created_at"), qm.Limit(10)).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// InsertReview _
func InsertReview(title, url, users string, groupID int64) {
	var review models.Review

	review.URL = url
	review.IsReviewed = false
	review.IsTested = false
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
	review.IsReviewed = false
	review.IsTested = false
	review.Title = title
	review.Users = users

	err = review.UpdateG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// UpdateToDoneReview _
func UpdateToDoneReview(sequences []string, groupID int64, user string, force bool) bool {
	reviews := GetAllNeedReview(groupID)
	successToUpdate := false

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

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

				successToUpdate = true
				break
			}
		}
	}

	return successToUpdate
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

func UpdateToDoneQA(sequences []string, groupID int64) bool {
	reviews := GetAllNeedQA(groupID)
	successToUpdate := false

	for _, seq := range sequences {
		sequence, err := strconv.Atoi(seq)
		if err != nil {
			continue
		}

		for i, review := range reviews {
			if i == sequence-1 {
				review.IsTested = true

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
