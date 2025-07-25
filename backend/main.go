package main

import (
	"backend/db"

	_ "github.com/go-sql-driver/mysql"
	"backend/routes"
)

func main() {
	// // Ganti ini dengan info koneksi database kamu
	// dsn := "root:@tcp(127.0.0.1:3306)/psc_db?parseTime=true"

	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }
	// defer db.Close()

	// // Registrasi route
	// router.RegisterUserRoutes(db)

	// // Start server
	// log.Println("Server running at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	db.Init()
	e := routes.
	e.Logger.Fatal(e.Start(":8080"))
}
