package app

//go:generate sqlboiler --wipe --add-global-variants --no-tests --no-context --no-rows-affected mysql

import (
	"bytes"
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/volatiletech/sqlboiler/boil"

	// Mysql
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Redis _
var Redis *redis.Client

// Bot _
var Bot *tb.Bot

var once sync.Once

// New _
func init() {
	// Load .env into terminal env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	once.Do(func() {
		resolveMysql()
		Redis = resolveRedis()
		Bot = resolveBot()
	})
}

func resolveMysql() {
	var b bytes.Buffer

	b.WriteString(os.Getenv("DATABASE_USERNAME"))
	b.WriteString(":")
	b.WriteString(os.Getenv("DATABASE_PASSWORD"))
	b.WriteString("@/")
	b.WriteString(os.Getenv("DATABASE_NAME"))
	b.WriteString("?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci")

	mysqlClient, err := sql.Open("mysql", b.String())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Mysql Loaded")

	boil.SetDB(mysqlClient)
}

func resolveBot() *tb.Bot {
	bot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %s", bot.Me.Username)

	return bot
}

func resolveRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Printf("Redis loaded")

	return client
}
