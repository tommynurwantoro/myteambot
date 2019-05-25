package mysql

import (
	"log"
	"strings"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/volatiletech/null"

	"github.com/bot/myteambot/app/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// InsertMessageRetro _
func InsertMessageRetro(username string, retroType string, message string) {
	var retro models.Retro

	retro.Username = username
	retro.Type = retroType
	retro.Message = null.StringFrom(message)

	err := retro.InsertG(boil.Infer())
	if err != nil {
		panic(err)
	}
}

// GetResultRetro _
func GetResultRetro(date string) []*models.Retro {
	splitDate := strings.Split(date, "-")
	recreatedDate := splitDate[2] + "-" + splitDate[1] + "-" + splitDate[0]

	retros, err := models.Retros(qm.Where("DATE(created_at) = ?", recreatedDate), qm.OrderBy("RAND()")).AllG()
	if err != nil {
		log.Fatal(err)
	}

	return retros
}
