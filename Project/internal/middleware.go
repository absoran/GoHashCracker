package internal

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// IsAuthorized function basically takes request and checks token is valid or not. If token is valid redirect to endpoint. If its not; returns unauthorized error
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwtToken")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
		}
		tokenstr := cookie.Value
		claims := &jwt.StandardClaims{}
		tkn, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if tkn.Valid {
			endpoint(w, r)
		}
	})
}
