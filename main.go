package main

import (
	"fmt"
	"log"
	"strconv"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
)

// Book struct represents a book
type Book struct {
	ID      int
	Judul   string
	Penulis string
	Rating  int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "go_nico"
	password = "12345"
	dbname   = "golangfiber"
)

// variable for slice
var books []Book

// Initialize the database connection
func initDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	} else {
		fmt.Println("success")
	}

	return db, nil
}

func main() {

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
		fmt.Println("ttt")
	}
	// defer db.Close()

	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		// Query data from the database
		rows, err := db.Query("SELECT id, judul, penulis, rating FROM books")
		if err != nil {
			log.Println("Error querying database:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to query database")
		}
		defer rows.Close()

		var books []Book
		for rows.Next() {
			var book Book
			if err := rows.Scan(&book.ID, &book.Judul, &book.Penulis, &book.Rating); err != nil {
				log.Println("Error scanning row:", err)
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to scan row")
			}
			books = append(books, book)
		}

		// Render a view with the retrieved data
		return c.Render("index", fiber.Map{
			"Books": books,
		})
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		// Extract form values
		judul := c.FormValue("judul")
		penulis := c.FormValue("penulis")
		rating, err := strconv.Atoi(c.FormValue("rating"))
		fmt.Println(judul, penulis, rating)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Rating")
		}

		// Insert the new book into the database
		_, err = db.Exec("INSERT INTO books (judul, penulis, rating) VALUES ($1, $2, $3)", judul, penulis, rating)
		if err != nil {
			log.Println("Error querying database:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to insert into database")
		}

		// Redirect to the home page
		return c.Redirect("/")
	})

	log.Fatal(app.Listen(":3002"))
}
