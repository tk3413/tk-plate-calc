package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	server "github.com/tk3413/tk-weight-calc/server_impl"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	srv := server.NewServer(server.WithLogger(logger))

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := server.HandlerFromMux(srv, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
