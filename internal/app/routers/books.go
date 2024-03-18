package routes

import (
	"github.com/caohoangphuctd97/go-test/internal/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// BookRoutes func for describe group of public routes.
func BookRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/books", controllers.GetBooks)   // get list of all books
	route.Get("/book/:id", controllers.GetBook) // get one book by ID

	// Routes for POST method:
	route.Post("/book", controllers.CreateBook) // create a new book

	// Routes for PATCH method:
	route.Patch("/book/:id", controllers.UpdateBook) // update one book by ID

	// Routes for DELETE method:
	route.Delete("/book/:id", controllers.DeleteBook) // delete one book by ID
}
