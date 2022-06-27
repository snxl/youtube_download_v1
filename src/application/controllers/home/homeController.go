package homeController

import "github.com/gofiber/fiber/v2"

func Handle(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "ok",
	})
}
