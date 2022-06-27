package homeController

import "github.com/gofiber/fiber/v2"

func Handle(ctx *fiber.Ctx) error {
	//return fiber.NewError(fiber.StatusBadRequest)
	return ctx.JSON(fiber.Map{
		"data": "ok",
	})
}
