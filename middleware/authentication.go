package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"wishlist/entities"
	"wishlist/helper"

	"github.com/dgrijalva/jwt-go"
)

type Access entities.Access

func CreateToken(uid string, email string, phone string, role string) (map[string]string, error) {
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

	return map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	}, nil
}

func RefreshToken() (string, error) {
	refresh := jwt.MapClaims{}
	refresh["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hour
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh)

	return tokenRefresh.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(r *http.Request) (*Access, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if !ok {
			return nil, err
		}
		uid := claims["uid"].(string)
		email := claims["email"].(string)
		phone := claims["phone"].(string)
		role := claims["role"].(string)
		return &Access{
			Uid:   uid,
			Email: email,
			Phone: phone,
			Role:  role,
		}, nil
	}
	return nil, err
}

func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if r.URL.Path == "/users/add" {
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
