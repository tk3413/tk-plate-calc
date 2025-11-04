package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	server "github.com/tk3413/tk-weight-calc/server_impl"
)

func main() {
	logger := setupLogger()
	srv := server.NewServer(server.WithLogger(logger))

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := server.HandlerFromMux(srv, r)
	s := &http.Server{
		Handler:           h,
		Addr:              "0.0.0.0:8080",
		ReadHeaderTimeout: 5 * time.Second,
	}

	logger.Info("Starting plate calc server", slog.String("addr", s.Addr))
	log.Fatal(s.ListenAndServe())
}

func setupLogger() *slog.Logger {
	level := strings.ToLower(strings.TrimSpace(os.Getenv("SLOG_LEVEL")))
	var handlerOpts *slog.HandlerOptions
	if level == "debug" {
		handlerOpts = &slog.HandlerOptions{Level: slog.LevelDebug}
	}
	return slog.New(slog.NewJSONHandler(os.Stderr, handlerOpts))
}
