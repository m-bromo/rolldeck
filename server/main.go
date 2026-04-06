package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/m-bromo/rolldeck/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	slog.Info("Starting application")

	mux := http.NewServeMux()

	http.ListenAndServe(cfg.Api.Addr, mux)
}
