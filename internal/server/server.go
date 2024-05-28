// Package server настройка и запуск сервиса
package server

import (
	"fmt"
	"loan/config"
	loancalc "loan/internal/loan_calc"
	"loan/internal/loan_calc/controller"
	"loan/internal/loan_calc/handler"
	"loan/internal/loan_calc/model"
	"loan/pkg/storage"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

// Server основной объект сервера, хранит необходимые для работы данные, в том числе конфиг, инстанс веб-сервера, ...
type Server struct {
	cfg *config.Config
	app *fiber.App
}

// New конструктор для Server.
func New(cfg *config.Config) *Server {
	app := fiber.New()
	return &Server{
		cfg: cfg,
		app: app,
	}
}

// Run запустить сервис.
func (s *Server) Run() {
	cache := storage.NewInMemCache[model.LoanResp]()
	loanController := controller.New(cache)
	loanHandler := handler.New(loanController)
	loancalc.SetupCreditRoutes(s.app, loanHandler)

	go s.listen()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := s.app.Shutdown(); err != nil {
		log.Println("failed to properly shutdown the server")
	}

	log.Println("app shutdown")
}

func (s *Server) listen() {
	addr := fmt.Sprintf(":%d", s.cfg.Server.Port)
	if err := s.app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
