package main

import (
	"encoding/json"
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
		t := time.Now()
		var out string
		if r.Header.Get("Accept") == "application/json" {
			out = buildJSON(t)
		} else {
			out = buildText(t)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(out))
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

func buildText(t time.Time) string {
	return t.Format(time.RFC3339)
}

func buildJSON(t time.Time) string {
	timeOut := struct {
		DayOfWeek  string `json:"day_of_week"`
		DayOfMonth int    `json:"day_of_month"`
		Month      string `json:"month"`
		Year       int    `json:"year"`
		Hour       int    `json:"hour"`
		Minute     int    `json:"minute"`
		Second     int    `json:"second"`
	}{
		DayOfWeek:  t.Weekday().String(),
		DayOfMonth: t.Day(),
		Month:      t.Month().String(),
		Year:       t.Year(),
		Hour:       t.Hour(),
		Minute:     t.Minute(),
		Second:     t.Second(),
	}
	out, _ := json.Marshal(timeOut)
	return string(out)
}
