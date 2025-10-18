package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	server := &http.Server{
		Addr:    ":4000",
		Handler: routes(),
	}

	logLevel := slog.LevelDebug
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

	logger.Info("Server is running on port :4000")

	err := server.ListenAndServe()

	if err != nil {
		os.Exit(1)
	}
}
