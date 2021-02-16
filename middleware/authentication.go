package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"
	"wishlist/helper"

	"github.com/dgrijalva/jwt-go"
)

type M map[string]string

func CreateToken(uid string, email string, phone string, role string) (M, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["uid"] = uid
	claims["email"] = email
	claims["phone"] = phone
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 3 hour

	access, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	refToken := jwt.New(jwt.SigningMethodHS256)
	refClaims := refToken.Claims.(jwt.MapClaims)
	refClaims["authorized"] = true
	refClaims["uid"] = uid
	refClaims["email"] = email
	refClaims["phone"] = phone
	refClaims["role"] = role
	refClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refresh, err := refToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return M{
		"access_token":  access,
		"refresh_token": refresh,
	}, nil
}

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if r.URL.Path == "/users/add" || r.URL.Path == "/login" {
            next.ServeHTTP(w, r)
            return
        }

		if tokenHeader == "" {
			resp := helper.Message(http.StatusUnauthorized, "Missing auth token")
			helper.Response(w, http.StatusUnauthorized, resp)
			return
		}

		splitted := strings.Split(tokenHeader, " ") 
		if len(splitted) != 2 {
			resp := helper.Message(http.StatusUnauthorized, "Missing auth token")
			helper.Response(w, http.StatusUnauthorized, resp)
			return
		}

		tokenPart := splitted[1] 

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			resp := helper.Message(http.StatusUnauthorized, "Missing auth token")
			helper.Response(w, http.StatusUnauthorized, resp)
			return
		}

		if !token.Valid {
			resp := helper.Message(http.StatusUnauthorized, "Missing auth token")
			helper.Response(w, http.StatusUnauthorized, resp)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
