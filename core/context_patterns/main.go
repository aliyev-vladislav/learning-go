package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func main() {
	ctx := context.Background()
	result, err := logic(ctx, "a string")
	fmt.Println(result, err)
}

func logic(ctx context.Context, info string) (string, error) {
	return "", nil
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		// wrap the context with stuff -- we'll see now soon!
		req = req.WithContext(ctx)
		handler.ServeHTTP(rw, req)
	})
}

func handler(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	data := req.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code %d", resp.StatusCode)
	}
	id, err := processResponse(resp.Body)
	return id, err
}

func processResponse(body io.ReadCloser) (string, error) {
	return "", nil
}
