package mysql

import (
	"log"
	"strings"

	"github.com/bot/act-bl-bot/app"
	"github.com/bot/act-bl-bot/entity"
)

// InsertMessageRetro _
func InsertMessageRetro(username string, retroType string, message string) {
	_, err := app.MysqlClient.Exec(
		"INSERT INTO retros(username, type, message, created_at, updated_at) VALUES(?, ?, ?, now(), now())",
		username, retroType, message)
	if err != nil {
		panic(err)
	}
}

// GetResultRetro _
func GetResultRetro(date string) []entity.Retro {
	//result := entity.Retro{}
	splitDate := strings.Split(date, "-")
	recreatedDate := splitDate[2] + "-" + splitDate[1] + "-" + splitDate[0]
	rows, err := app.MysqlClient.Query("SELECT * FROM retros WHERE DATE(created_at) = ? ORDER BY RAND()", recreatedDate)
	if err != nil {
		log.Fatal(err)
	}

	results := make([]entity.Retro, 0)
	for rows.Next() {
		var result entity.Retro
		if err := rows.Scan(&result.ID, &result.Username, &result.Type, &result.Message, &result.CreatedAt, &result.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	return results
}
