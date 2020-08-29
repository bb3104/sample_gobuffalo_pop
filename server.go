package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gobuffalo/pop"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wapa5pow/go-database-library-how-to/pop/models"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := pop.Connect("development")
	if err != nil {
		log.Fatal(err)
	}
	users := []*models.User{}

	err = db.All(&users)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", users)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", helloWorld())

	err = e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}

func helloWorld() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "helloWorld!")
	}
}
