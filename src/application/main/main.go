package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/snxl/youtube-dowloader/src/application/middlewares"
	"github.com/snxl/youtube-dowloader/src/application/routes"
	"log"
	"time"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler,
		Prefork:      true,
		Immutable:    true,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	loadMiddlewares(app)
	loadRoutes(app)

	log.Fatal(app.Listen(":8000"))
}

func loadMiddlewares(app *fiber.App) {
	app.Use(recover.New())
	app.Use(favicon.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Next: middlewares.ProcessesChild,
	}))
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Second,
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/favicon.ico"
		},
	}))
}

func loadRoutes(app *fiber.App) {
	routes.SetupRoutes(app)
}
