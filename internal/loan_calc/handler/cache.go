package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Cache хендлер для GET /cache.
func (h *Handler) Cache(c *fiber.Ctx) error {
	records, err := h.ctrl.List(c.Context())
	if err != nil {
		c = c.Status(http.StatusBadRequest)
		return fiber.NewError(c.Response().StatusCode(), err.Error())
	}
	if err := c.Status(http.StatusOK).JSON(records); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
