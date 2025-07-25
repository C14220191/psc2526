package main

import (
	"backend/db"

	"backend/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	conn := db.CreateCon()
	e := echo.New()
	routes.RegisterUserRoutes(e, conn)
	routes.RegisterAssessmentRoutes(e, conn)

	e.Logger.Fatal(e.Start(":8080"))
}
