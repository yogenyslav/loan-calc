// Package loancalc корневой пакет домена сервиса
package loancalc

import (
	"loan/internal/loan_calc/middleware"

	"github.com/gofiber/fiber/v2"
)

type creditHandler interface {
	Execute(c *fiber.Ctx) error
	Cache(c *fiber.Ctx) error
}

// SetupCreditRoutes маппинг путей с хендлерами.
func SetupCreditRoutes(app *fiber.App, h creditHandler) {
	g := app.Group("/api")

	g.Use(middleware.Logger())
	g.Post("/execute", h.Execute)
	g.Get("/cache", h.Cache)
}
