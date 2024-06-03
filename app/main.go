package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

func main() {
	connctionString := fmt.Sprintf("%s:%s@tcp(mariadb:3306)/madb", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"))
	if os.Getenv("MYSQL_USER") == "" {
		panic("MYSQL_USER is not set")
	}
	if os.Getenv("MYSQL_PASSWORD") == "" {
		panic("MYSQL_PASSWORD is not set")
	}
	if os.Getenv("AUTH") == "" {
		panic("AUTH is not set")
	}
	if os.Getenv("TOKEN") == "" {
		panic("TOKEN is not set")
	}

	fmt.Println("Connecting to", connctionString)
	var err error
	db, err = sql.Open("mysql", connctionString)
	if err != nil {

		fmt.Println("Error: Could not establish a connection with the database", err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error: Could not establish a connection with the database", err)
		os.Exit(1)
	}

	defer db.Close()

	e := echo.New()
	e.GET("/w/:wire", redir)
	e.POST("/api/addwire", addWire)
	e.POST("/api/getlogs", getLogs)
	e.File("/addwire", "public/addwire.html")
	e.File("/auth", "public/auth.html")
	e.File("/cam", "public/cam.html")
	e.POST("/api/directlog", directLog)

	e.Logger.Fatal(e.Start(":3000"))
}
