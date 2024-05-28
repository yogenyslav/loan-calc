package handler

import (
	"loan/internal/loan_calc/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Execute хендлер для POST /execute, ожидает model.LoanReq в теле запроса.
func (h *Handler) Execute(c *fiber.Ctx) error {
	var req model.LoanReq
	if err := c.BodyParser(&req); err != nil {
		c = c.Status(http.StatusUnprocessableEntity)
		return fiber.NewError(c.Response().StatusCode(), err.Error())
	}

	resp, err := h.ctrl.Execute(c.Context(), req)
	if err != nil {
		c = c.Status(http.StatusBadRequest)
		return fiber.NewError(c.Response().StatusCode(), err.Error())
	}

	if err := c.Status(http.StatusOK).JSON(struct {
		Result model.LoanResp `json:"result"`
	}{resp}); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
