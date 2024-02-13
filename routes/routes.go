package routes

import (
	"gofiber/handler"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App, bookHandler *handler.BookHandler) {
	app.Static("/", "./static")
	app.Static("/edit", "./static")

	app.Get("/", bookHandler.GetBooks)
	app.Post("/add", bookHandler.AddBook)
	app.Post("/delete/:id", bookHandler.DeleteBook)
	app.Get("/edit/:id", bookHandler.EditBook)
	app.Post("/update/:id", bookHandler.UpdateBook)
}
