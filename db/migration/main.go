package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"

	// Mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load .env into terminal env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migration/mysql",
	}

	var b bytes.Buffer
	b.WriteString(os.Getenv("DATABASE_USERNAME"))
	b.WriteString(":")
	b.WriteString(os.Getenv("DATABASE_PASSWORD"))
	b.WriteString("@/")
	b.WriteString(os.Getenv("DATABASE_NAME"))
	b.WriteString("?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci")

	mysqlClient, err := sql.Open("mysql", b.String())
	if err != nil {
		panic(err)
	}

	defer mysqlClient.Close()

	n, err := migrate.Exec(mysqlClient, "mysql", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
