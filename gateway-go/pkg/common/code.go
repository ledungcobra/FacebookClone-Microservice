package common

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateCode(length int) string {
	rand.Seed(time.Now().Unix())

	randomCode := ""
	for i := 0; i < length; i++ {
		randomCode += strconv.Itoa(rand.Intn(10))

	}
	return randomCode
}
