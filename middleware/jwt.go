package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Filter func(c *fiber.Ctx) bool
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Filter != nil && config.Filter(c) {
			return c.Next()
		}

		authHeader := c.Get("Authorization")

		fmt.Println(authHeader)

		return c.Next()
	}
}
