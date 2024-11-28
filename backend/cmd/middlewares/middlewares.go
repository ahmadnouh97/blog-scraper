package middlewares

import "net/http"

func ApiKeyMiddleware(next http.Handler, validAPIKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get API key from the `Authorization` header
		apiKey := r.Header.Get("Authorization")

		// Validate the API key
		if apiKey != validAPIKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
