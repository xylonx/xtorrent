package utils

import (
	"backend/env"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"time"
)

var (
	JwtConfig = middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		ContextKey:    "user",
		Claims:        &JwtClaim{},
		SigningKey:    []byte(env.Conf.Jwt.Secret),
	}
)

type JwtClaim struct {
	jwt.StandardClaims
}

func GenerateJwtToken(id string) (string, error) {
	claims := &JwtClaim{
		jwt.StandardClaims{
			Id:        id,
			ExpiresAt: time.Now().Add(time.Duration(env.Conf.Jwt.ExpireDuration)).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(env.Conf.Jwt.Secret))
	if err != nil {
		return "", err
	}

	log.Println("generate token for id: " + id)

	return token, nil
}

func ValidateJwtToken(context echo.Context) bool {

	claims := getClaim(context)

	if claims.ExpiresAt > time.Now().Unix() {
		return false
	}

	return true
}

func GetCurrentUserEmail(context echo.Context) string {
	return getClaim(context).Id
}

func getClaim(context echo.Context) *JwtClaim {

	user := context.Get("user").(*jwt.Token)
	return user.Claims.(*JwtClaim)
}
