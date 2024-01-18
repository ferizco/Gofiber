package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// Book struct represents a book
type Book struct {
	Judul   string
	Penulis string
	Rating  int
}

// variable for slice
var books []Book

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Books": books,
		})
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		// Extract form values
		judul := c.FormValue("judul")
		penulis := c.FormValue("penulis")
		rating, err := strconv.Atoi(c.FormValue("rating")) // cannot use int as value in struct
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Rating")
		}

		// Create a new book
		book := Book{
			Judul:   judul,
			Penulis: penulis,
			Rating:  rating,
		}

		// Append the new book to the slice
		books = append(books, book)

		// Redirect to the home page
		return c.Redirect("/")
	})

	log.Fatal(app.Listen(":3002"))
}
