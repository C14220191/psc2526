package main

import (
	"backend/router"
	"database/sql"
	"log"
	

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/psc_db")
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()

	e := echo.New()
	router.RegisterUserRoutes(e, db)

	log.Println("Server running at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
