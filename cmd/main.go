package main

import (
	"log"
	"os"

	"github.com/dattruongdev/bookstore_cqrs/app"
	"github.com/dattruongdev/bookstore_cqrs/config"
	"github.com/dattruongdev/bookstore_cqrs/database"
	"github.com/dattruongdev/bookstore_cqrs/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.NewConfig()
	log.Println("This is conf", conf)
	db := database.Connect(conf)
	server_port := os.Getenv("SERVER_PORT")

	e := echo.New()

	app := app.NewApplication(db)

	route.AddRoutes(e, db, &app)

	e.Logger.Fatal(e.Start(server_port))
}
