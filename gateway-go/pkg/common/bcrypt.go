package common

import (
	"ledungcobra/gateway-go/pkg/config"

	"golang.org/x/crypto/bcrypt"
)

func init() {

}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.Cfg.GatewayCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RecoverPassword(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) != nil
}
