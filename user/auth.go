package user

import (
	"errors"
	"net/http"
	"qk-note/shared"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Auth : Authentication Middleware
type Auth interface {
	VerifyToken(http.Handler) http.Handler
	GenerateToken(user *Claims) (*JWT, error)
}

// JWTAuth : JWT Auth Struct
type JWTAuth struct{}

// NewAuthMiddleware : Returns Auth Struct
func NewAuthMiddleware() Auth {
	return &JWTAuth{}
}

// GenerateToken : to Generate JWT Access Token
func (a *JWTAuth) GenerateToken(user *Claims) (*JWT, error) {
	var jwtKey = []byte(shared.JWTKey)
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return nil, errors.New(shared.JWTError)
	}
	var jwtToken JWT
	jwtToken.Token = token
	return &jwtToken, nil
}

// VerifyToken : to Verify JWT Token
func (a *JWTAuth) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		authHeader := strings.Split(bearerToken, " ")
		if len(authHeader) != 2 {
			shared.Fail(w, 400, shared.InvalidAccessTokenErrorCode, shared.InvalidAccessTokenError)
			return
		}
		jwtToken := authHeader[1]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(shared.JWTKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				shared.Fail(w, 400, shared.InvalidAccessTokenErrorCode, shared.InvalidAccessTokenError)
				return
			}
			shared.Fail(w, 400, shared.InvalidAccessTokenErrorCode, shared.InvalidAccessTokenError)
			return
		}
		if !token.Valid {
			shared.Fail(w, 400, shared.InvalidAccessTokenErrorCode, shared.InvalidAccessTokenError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
