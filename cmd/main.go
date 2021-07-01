package main

import (
		"eventlog/cmd/db"
		"eventlog/cmd/routes"
		"eventlog/infrastructure/repository"
		"eventlog/service/event"
		"fmt"
		"github.com/gofiber/fiber/v2"
		_ "github.com/lib/pq"
		"log"
		"os"
)

func main() {

		initApp()
		eventPgRepo := repository.NewRepository(db.DB)
		eventService := event.NewService(eventPgRepo)
		app := fiber.New()
		api := app.Group("/api")
		routes.Routes(api, eventService)
		log.Fatal(app.Listen(":3000"))
}

func initApp() {

		var (
				user     = os.Getenv("POSTGRES_USER")
				password = os.Getenv("POSTGRES_PASSWORD")
				dbname   = os.Getenv("POSTGRES_DB")
				host     = os.Getenv("POSTGRES_HOST")
				port     = os.Getenv("POSTGRES_PORT")
		)

		// Init db
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
				"password=%s dbname=%s sslmode=disable",
				host, port, user, password, dbname)
		db.InitDB(psqlInfo)

}
