package main

import (
	"github.com/api_go_opp/router"
	"golang.org/x/exp/slog"
)

func main() {
	slog.Info("Iniciando a aplicação...")
	router.Init()
}
