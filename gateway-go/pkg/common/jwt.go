package common

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JSON = map[string]interface{}

func String(j JSON) string {
	v, _ := json.MarshalIndent(j, "", " ")
	return string(v)
}

func GenerateToken(data JSON, expired time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, createClaim(expired, data))
	tokenString, err := token.SignedString([]byte(os.Getenv("GATEWAY_JWT_SECRET")))
	return tokenString, err
}

func createClaim(expired time.Duration, data JSON) jwt.Claims {
	return jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(expired)},
		Issuer:    "gateway",
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		Subject:   String(data),
	}
}

func extractTokenFromRequest(request *fasthttp.Request) string {
	bearToken := string(request.Header.Peek("Authorization"))
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func extractToken(inputToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(inputToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("GATEWAY_JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractClaim(request *fasthttp.Request) (*jwt.RegisteredClaims, error) {
	var err error
	tokenStr := extractTokenFromRequest(request)
	if tokenStr == "" {
		return nil, fmt.Errorf("token is empty")
	}
	var token *jwt.Token
	if token, err = extractToken(tokenStr); err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("claim is not registered claim")
	}
	return claim, nil
}

func ExtractFromString(tokenStr string) (*jwt.RegisteredClaims, error) {
	var err error
	if tokenStr == "" {
		return nil, fmt.Errorf("token is empty")
	}
	var token *jwt.Token
	if token, err = extractToken(tokenStr); err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("claim is not registered claim")
	}
	return claim, nil
}
