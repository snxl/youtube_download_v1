package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func ProcessesChild(c *fiber.Ctx) bool {
	return !fiber.IsChild()
}
