package jwtgenerator

import (
	"fmt"
	"time"

	"github.com/erupshis/revtracker/internal/logger"
	"github.com/golang-jwt/jwt/v4"
)

// Claims struct that keeps standard jwt claims plus custom UserID.
type Claims struct {
	jwt.RegisteredClaims
	UserID int64
}

// JwtGenerator generator itself.
type JwtGenerator struct {
	jwtKey   string
	tokenExp int

	log logger.BaseLogger
}

// Create creates JWT tokens generator.
func Create(jwtKey string, tokenExp int, baseLogger logger.BaseLogger) JwtGenerator {
	if jwtKey == "" {
		baseLogger.Info("[auth:jwtgenerator:Create] JWT token generation key is missing")
	}

	return JwtGenerator{
		jwtKey:   jwtKey,
		tokenExp: tokenExp,
		log:      baseLogger,
	}
}

// BuildJWTString creates token and returns it as string.
func (j *JwtGenerator) BuildJWTString(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.tokenExp) * time.Hour)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(j.jwtKey))
	if err != nil {
		return "", err
	}

	j.log.Info("[auth:jwtgenerator:BuildJWTString] created JWT token for userID '%d'", userID)
	return tokenString, nil
}

// GetUserID gets token in string format, parse it and returns userID.
func (j *JwtGenerator) GetUserID(tokenString string) int64 {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(j.jwtKey), nil
		})
	if err != nil {
		return -1
	}

	if !token.Valid {
		j.log.Info("[auth:jwtgenerator:GetUserID] Token is not valid")
		return -1
	}

	j.log.Info("[auth:jwtgenerator:GetUserID] Token os valid")
	return claims.UserID
}
