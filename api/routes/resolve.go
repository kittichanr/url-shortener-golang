package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kittichanr/url-shortener-golang/database"
	"github.com/redis/go-redis/v9"
)

func Resolve(ctx *fiber.Ctx) error {
	url := ctx.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short-url not found in db"})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal error"})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return ctx.Redirect(value, 301)
}
