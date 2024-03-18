package main

import (
	"os"

	routes "github.com/caohoangphuctd97/go-test/internal/app/routers"
	"github.com/caohoangphuctd97/go-test/pkg/configs"
	middleware "github.com/caohoangphuctd97/go-test/pkg/middlewares"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title GO exercise #2
// @version 1.0
// @description This is a swagger for go exercise #2.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	zerolog.TimeFieldFormat = zerolog.TimeFieldFormat

	log.Info().Msg("Start server")
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app) // Register a swagger APIs
	routes.BookRoutes(app)   // Register a public routes for app.

	app.Listen("0.0.0.0:8080")
	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
