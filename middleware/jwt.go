package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/swildz/go-jwt-siddiq/config"
	"github.com/swildz/go-jwt-siddiq/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {

				response := map[string]string{"Message ": "Unauthorized"}
				helper.ResponJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		//mengambil token string

		tokenString := c.Value

		claims := &config.JWTClaim{}

		//parsing token jwt

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)

			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				//token invalid
				response := map[string]string{"Message ": "Unauthorized"}
				helper.ResponJSON(w, http.StatusUnauthorized, response)
				return

			case jwt.ValidationErrorExpired:
				//token expaired
				response := map[string]string{"Message ": "Unauthorized, Token Expaired"}
				helper.ResponJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"Message ": "Unauthorized"}
				helper.ResponJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"Message ": "Unauthorized"}
			helper.ResponJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)

	})

}
