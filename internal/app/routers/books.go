package routes

import (
	"github.com/caohoangphuctd97/go-test/internal/app/controllers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type BookCntrlImpl struct {
	dig.In
	Svc controllers.BookSvc
}

type BookRoutes interface {
	SetRoute(a *fiber.App)
}

func NewBookCntrl(impl BookCntrlImpl) BookRoutes {
	return &impl
}

// BookRoutes func for describe group of public routes.
func (c *BookCntrlImpl) SetRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/books", c.Svc.GetBooks)   // get list of all books
	route.Get("/book/:id", c.Svc.GetBook) // get one book by ID

	// Routes for POST method:
	route.Post("/book", c.Svc.CreateBook) // create a new book

	// Routes for PATCH method:
	route.Patch("/book/:id", c.Svc.UpdateBook) // update one book by ID

	// Routes for DELETE method:
	route.Delete("/book/:id", c.Svc.DeleteBook) // delete one book by ID
}
