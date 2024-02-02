package usecase

import (
	"gofiber/models"
	"gofiber/repository"
)

type BookUsecase interface {
	GetBooks() ([]models.Book, error)
	AddBook(judul, penulis, rating string) error
}

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (uc *bookUsecase) GetBooks() ([]models.Book, error) {
	return uc.bookRepo.GetAll()
}

func (uc *bookUsecase) AddBook(judul, penulis, rating string) error {
	book := models.Book{
		Judul:   judul,
		Penulis: penulis,
		Rating:  rating,
	}

	return uc.bookRepo.Add(book)
}
