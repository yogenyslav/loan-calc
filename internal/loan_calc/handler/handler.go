// Package handler эндпоинты сервиса, обработка входных данных.
package handler

import (
	"context"
	"loan/internal/loan_calc/model"
)

type loanController interface {
	Execute(ctx context.Context, params model.LoanReq) (model.LoanResp, error)
	List(ctx context.Context) ([]model.CacheResp, error)
}

// Handler основной объект-обработчик, методами которого реализуют обработку входных данных и вызов методов контроллера для соответствующих ручек.
type Handler struct {
	ctrl loanController
}

// New конструктор для Handler.
func New(controller loanController) *Handler {
	return &Handler{
		ctrl: controller,
	}
}
