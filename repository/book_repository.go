package repository

import (
	"database/sql"
	"fmt"
	"gofiber/models"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	Add(models.Book) error
	DeleteBookByID(id int) error
	GetBookByID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
}

type PostgreSQLRepository struct {
	db *sql.DB
}

func NewPostgreSQLRepository(db *sql.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		db: db,
	}
}

func (r *PostgreSQLRepository) GetAll() ([]models.Book, error) {
	rows, err := r.db.Query("SELECT id, judul, penulis, rating FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Judul, &book.Penulis, &book.Rating); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *PostgreSQLRepository) Add(b models.Book) error {
	_, err := r.db.Exec("INSERT INTO books (judul, penulis, rating) VALUES ($1, $2, $3)", b.Judul, b.Penulis, b.Rating)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLRepository) DeleteBookByID(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLRepository) GetBookByID(id int) (*models.Book, error) {
	row := r.db.QueryRow("SELECT id, judul, penulis, rating FROM books WHERE id = $1", id)

	var book models.Book

	err := row.Scan(&book.ID, &book.Judul, &book.Penulis, &book.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no book found with ID %d", id)
		}
		return nil, err
	}

	return &book, nil
}

func (r *PostgreSQLRepository) UpdateBook(book *models.Book) error {

	_, err := r.GetBookByID(book.ID)
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	query := "UPDATE books SET judul=$1, penulis=$2, rating=$3 WHERE id=$4"
	_, err = r.db.Exec(query, book.Judul, book.Penulis, book.Rating, book.ID)

	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	return nil
}
