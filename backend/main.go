package main

import (
	"backend/db"
	"backend/validator"

	"backend/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	conn := db.CreateCon()
	validator.InitValidator()
	e := echo.New()
	e.Validator = validator.GetEchoAdapter()
	routes.RegisterUserRoutes(e, conn)
	routes.RegisterAssessmentRoutes(e, conn)
	routes.AdminRoute(e, conn)

	e.Logger.Fatal(e.Start(":8080"))
}
