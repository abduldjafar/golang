package auth

import (
	"backend-code/model"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//Exception struct
type Exception model.Exception

// JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "Missing auth token"})
			return
		}
		tk := &model.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", "")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//this code copy from https://raw.githubusercontent.com/linux08/golang-jwt-authentication/master/utils/auth/auth.go
