package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type ContextKey string

const AuthenticationContextKey ContextKey = "auth"

type Claims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	jwt.RegisteredClaims
}

type Authentication struct {
	jwtKey         []byte
	expirationTime time.Duration
}

func AuthenticationFactory(jwtKey string, expirationTime time.Duration) Authentication {
	return Authentication{
		jwtKey:         []byte(jwtKey),
		expirationTime: expirationTime,
	}
}

func (a *Authentication) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := a.extractToken(r)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		err = a.validateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *Authentication) AuthenticationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), AuthenticationContextKey, a)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Authentication) InitJwtKey(key string) {
	a.jwtKey = []byte(key)
}

func (a *Authentication) GenerateToken(username string, id string) (string, error) {
	expirationTime := time.Now().Add(a.expirationTime)
	claims := &Claims{
		Username: username,
		Id:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(a.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Format "Bearer <token>".
func (a *Authentication) extractToken(r *http.Request) (string, error) {
	authenticationHeader := r.Header.Get("Authorization")
	if authenticationHeader == "" {
		return "", errors.New("Authentication header not found")
	}

	tokenString := strings.TrimPrefix(authenticationHeader, "Bearer ")
	if tokenString == authenticationHeader {
		return "", errors.New("Invalid authentication header format")
	}

	return tokenString, nil
}

func (a *Authentication) validateToken(tokenString string) error {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return a.jwtKey, nil
		},
	)

	if err != nil || !token.Valid {
		return errors.New("Invalid token")
	}
	return nil
}

// ==================================================Password authentication==================================================

func (a *Authentication) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (a *Authentication) ValidatePassword(password string, passwordHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}

// Others
func GetAuthFromContext(r *http.Request) *Authentication {
	auth, _ := r.Context().Value(AuthenticationContextKey).(*Authentication)
	return auth
}
