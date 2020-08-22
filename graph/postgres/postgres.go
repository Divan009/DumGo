package postgres

import (
	sql "database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.

	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "@Divas009"
	dbname := "postgres"
	sslmode := "disable"

	PsqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode

	db, err = sql.Open("postgres", PsqlInfo)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Postgres connected!")

}

func GetDb() *sql.DB {
	if db == nil {
		fmt.Println("Reconnecting to Postgres DB!")
		InitDB()
	}
	return db
}
