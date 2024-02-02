package handler

import (
	"gofiber/usecase"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	bookUsecase usecase.BookUsecase
}

func NewBookHandler(bookUsecase usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bookUsecase: bookUsecase,
	}
}

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.bookUsecase.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve books")
	}

	return c.Render("index", fiber.Map{
		"Books": books,
	})
}

func (h *BookHandler) AddBook(c *fiber.Ctx) error {
	judul := c.FormValue("judul")
	penulis := c.FormValue("penulis")
	rating := c.FormValue("rating")

	err := h.bookUsecase.AddBook(judul, penulis, rating)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add a new book")
	}

	return c.Redirect("/")
}
