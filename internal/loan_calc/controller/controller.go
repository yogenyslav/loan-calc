// Package controller бизнес-логика всех эндпоинтов.
package controller

import (
	"loan/internal/loan_calc/model"
	"loan/pkg/storage"
)

// Controller основной объект-обработчик, реализующий бизнес-логику сервиса.
type Controller struct {
	cache *storage.InMemCache[model.LoanResp]
}

// New конструктор для Controller.
func New(cache *storage.InMemCache[model.LoanResp]) *Controller {
	return &Controller{
		cache: cache,
	}
}
