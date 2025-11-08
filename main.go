package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	logLevel := slog.LevelDebug
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

	app := &application{
		logger: logger,
	}

	server := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	logger.Info("The server is running on port :4000")
	err := server.ListenAndServe()

	if err != nil {
		os.Exit(1)
	}
}
