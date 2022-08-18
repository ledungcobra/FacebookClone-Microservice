package common

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func init() {

}

func HashPassword(password string) (string, error) {
	cost, _ := strconv.Atoi(os.Getenv("GATEWAY_BCRYPT_COST"))
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RecoverPassword(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) != nil
}
