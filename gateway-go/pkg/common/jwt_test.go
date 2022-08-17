package common

import (
	"sync"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type args struct {
	username string
	expired  time.Duration
}

type input struct {
	name string
	args args
}

func TestJWTTokenValid(t *testing.T) {

	tests := []input{
		{
			name: "test generate token",
			args: args{username: "ledungcobra", expired: time.Second * 5},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		go func(tt input) {
			defer wg.Done()
			generatedToken, err := GenerateToken(JSON{"username": tt.args.username}, tt.args.expired)
			t.Logf("Generated token %s", generatedToken)
			if err != nil {
				t.Error(err)
				return
			}

			time.Sleep(tt.args.expired - 2*time.Second)
			token, err := VerifyToken(generatedToken)
			if err != nil {
				t.Error("verify token error")
				return
			}
			if !token.Valid {
				t.Error("token is not valid")
			}
			userName := token.Claims.(jwt.MapClaims)["username"].(string)
			if tt.args.username != userName {
				t.Errorf("%s: want username %s, got %s", tt.name, tt.args.username, userName)
				return
			}
		}(tt)
	}
	wg.Wait()
}

func TestJWTTokenInvalid(t *testing.T) {

	tests := []input{
		{
			name: "test generate token",
			args: args{username: "ledungcobra", expired: time.Second},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		go func(tt input) {
			defer wg.Done()
			generatedToken, err := GenerateToken(JSON{"username": tt.args.username}, tt.args.expired)
			if err != nil {
				t.Error(err)
			}

			time.Sleep(tt.args.expired + 2*time.Second)
			token, err := VerifyToken(generatedToken)
			if err == nil {
				t.Error("Should return error in case of token invalid")
			}
			if token != nil {
				t.Error("Token should be nil when invalid")
			}
		}(tt)
	}
	wg.Wait()
}
