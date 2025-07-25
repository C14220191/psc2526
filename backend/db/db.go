package db

import (
	"backend/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?parseTime=true"
	fmt.Println("Connection string:", connectionString)
	db, err = sql.Open("mysql", connectionString)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Second * 5)


	if err != nil {
		panic("connectionString error...")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN invalid")
	}
}
func CreateCon() *sql.DB {
	return db
}
func Close() {
	if db != nil {
		err = db.Close()
		if err != nil {
			fmt.Println("Error closing database connection:", err)
		}
	}
}
