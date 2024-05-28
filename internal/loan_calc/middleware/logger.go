// Package middleware хранит кастомные middleware
package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Logger middleware для логирования статуса и времени выполнения запроса.
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		log.Printf("status_code: %d, duration: %d ns", c.Response().StatusCode(), time.Since(start).Nanoseconds())
		if err != nil {
			return fiber.NewError(c.Response().StatusCode(), err.Error())
		}
		return nil
	}
}
