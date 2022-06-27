package routes

import (
	"github.com/gofiber/fiber/v2"
	homeController "github.com/snxl/youtube-dowloader/src/application/controllers/home"
)

func loadMainRoutes(app *fiber.App) {
	app.Get("/", homeController.Handle)
}
