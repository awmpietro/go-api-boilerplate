package middleware

import (
	"context"
	"net/http"
	"os"

	"teste-go/internal/entity"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const (
	userContextKey contextKey = "user"
)

// JWTMiddleware is a middleware that verifies JWT tokens
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the JWT token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse and verify the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if token.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrSignatureInvalid
			}
			secretKey := []byte(os.Getenv("SECRET_KEY"))
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusBadRequest)
			return
		}
		email, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "Invalid email claim", http.StatusBadRequest)
			return
		}
		id, ok := claims["id"].(int)
		if !ok {
			http.Error(w, "Invalid id claim", http.StatusBadRequest)
			return
		}
		var user = &entity.User{
			ID:    id,
			Email: &email,
		}

		// Set the authenticated user in the request context or add any additional logic
		// Example: r = r.WithContext(context.WithValue(r.Context(), "user", user))
		ctx := context.WithValue(r.Context(), userContextKey, user)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}
