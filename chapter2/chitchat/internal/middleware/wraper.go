package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func WraperMiddleware(handler http.Handler, middleware ...Middleware) http.Handler {
	for _, mw := range middleware {
		handler = mw(handler)
	}

	return handler
}
