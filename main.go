package main

import (
	postgres "gofiber/database"
	"gofiber/handler"
	"gofiber/repository"
	"gofiber/routes"
	"gofiber/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})

	// Initialize database connection
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize dependencies
	bookRepo := repository.NewPostgreSQLRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	// Set up routes using the routes package
	routes.SetRoutes(app, bookHandler)

	// Start the server
	log.Fatal(app.Listen(":3002"))
}
