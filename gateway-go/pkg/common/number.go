package common

import (
	"log"
	"strconv"
)

func ToInt(tagFieldName string) int {
	val, err := strconv.Atoi(tagFieldName)
	if err != nil {
		log.Print("Error ", err.Error())
	}
	return val
}
