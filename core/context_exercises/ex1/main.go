package main

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	middleware := Timeout(100000)
	server := http.Server{
		Handler: middleware(http.HandlerFunc(sleepy)),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}

func Timeout(ms int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(ms))
			defer cancelFunc()
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}

}

func sleepy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	message, err := doThing(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write([]byte(message))
}

func doThing(ctx context.Context) (string, error) {
	wait := rand.Intn(5)
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		return "Done!", nil
	case <-ctx.Done():
		return "Too slow!", ctx.Err()
	}
}
