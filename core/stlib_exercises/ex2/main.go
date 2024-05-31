package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	options := &slog.HandlerOptions{}
	slogger := slog.New(slog.NewJSONHandler(os.Stderr, options))
	mux := createServeMux(slogger)
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}

func createServeMux(logger *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format(time.RFC3339)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(t))
	})
	mux.Handle("GET /", middlewareLogger(finalHandler, logger))
	return mux
}

func middlewareLogger(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ip, _, _ := strings.Cut(req.RemoteAddr, ":")
		logger.Info("incoming IP", "ip", ip)
		next.ServeHTTP(w, req)
	})
}
