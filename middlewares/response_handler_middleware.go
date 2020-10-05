package middlewares

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("custom-header", "this is my custom header")

		next.ServeHTTP(w, req)
	})
}
