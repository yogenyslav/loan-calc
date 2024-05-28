// Package main точка входа сервиса
package main

import (
	"loan/config"
	"loan/internal/server"
)

func main() {
	cfg := config.MustNew("./config/config.yml")

	srv := server.New(cfg)
	srv.Run()
}
