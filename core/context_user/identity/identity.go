package identity

import (
	"context"
	"net/http"
)

type userKey int

const (
	_ userKey = iota
	key
)

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

func userFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("identity")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		user, err := extractUser(req)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("unauthorized"))
			return
		}
		ctx := req.Context()
		ctx = ContextWithUser(ctx, user)
		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}

func SetUser(user string, rw http.ResponseWriter) {
	http.SetCookie(rw, &http.Cookie{
		Name:  "identity",
		Value: user,
	})
}

func DeleteCookie(rw http.ResponseWriter) {
	http.SetCookie(rw, &http.Cookie{
		Name:   "identity",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
