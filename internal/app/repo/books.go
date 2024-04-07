package repo

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"go.uber.org/dig"

	sq "github.com/Masterminds/squirrel"
)

// Book struct to describe book object.
type (
	Book struct {
		ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
		CreatedAt time.Time `db:"created_at" json:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
		Title     string    `db:"title" json:"title" validate:"required,lte=255"`
		Author    string    `db:"author" json:"author" validate:"required,lte=255"`
	}
	BookRepo interface {
		GetBooks() ([]Book, error)
		GetBook(uuid.UUID) (Book, error)
		CreateBook(*Book) error
		UpdateBook(uuid.UUID, *Book) error
		DeleteBook(uuid.UUID) error
	}
	BookRepoImpl struct {
		dig.In
		*sql.DB `name:"pg"`
	}
)

func NewBookRepo(impl BookRepoImpl) BookRepo {
	return &impl
}

// GetBooks method for getting all books.
func (q *BookRepoImpl) GetBooks() ([]Book, error) {
	// Define books variable.
	books := []Book{}

	rows, err := sq.Select("*").From("books").RunWith(q.DB).Query()
	if err != nil {
		// Return empty object and error.
		return books, err
	}

	for rows.Next() {
		ent := Book{}
		if err = rows.Scan(
			&ent.ID,
			&ent.Title,
			&ent.Author,
			&ent.UpdatedAt,
			&ent.CreatedAt,
		); err != nil {
			return books, err
		}
		books = append(books, ent)
	}
	// Return query result.
	return books, nil
}

// GetBook method for getting one book by given ID.
func (q *BookRepoImpl) GetBook(id uuid.UUID) (Book, error) {
	// Define book variable.
	book := Book{}

	rows, err := sq.Select("*").From("books").Where(sq.Eq{"id": id}).RunWith(q.DB).Query()
	if err != nil {
		// Return empty object and error.
		return book, err
	}
	if err = rows.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.UpdatedAt,
		&book.CreatedAt,
	); err != nil {
		return book, err
	}

	// Return query result.
	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *BookRepoImpl) CreateBook(b *Book) error {
	// Define query string.
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.Title, b.Author, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateBook method for updating book by given Book object.
func (q *BookRepoImpl) UpdateBook(id uuid.UUID, b *Book) error {
	// Define query string.
	query := `UPDATE books SET updated_at = $2, title = $3, author = $4 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id, b.UpdatedAt, b.Title, b.Author)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteBook method for delete book by given ID.
func (q *BookRepoImpl) DeleteBook(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
