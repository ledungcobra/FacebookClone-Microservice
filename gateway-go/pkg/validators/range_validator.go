package validators

import (
	"fmt"
	"log"
)

type RangeValidator struct {
	min    int
	max    int
	number int
}

func (m RangeValidator) GetErrorMessage() string {
	return fmt.Sprintf("Field must between %d and %d", m.min, m.max)
}

func (m RangeValidator) Validate() bool {
	return m.min <= m.number && m.number <= m.max
}

func NewRangeValidator(number int, min int, max int) IValidator {
	if min > max {
		log.Panicf("Min = %d, max %d is invalid input", min, max)
	}
	return &RangeValidator{
		min:    min,
		number: number,
		max:    max,
	}
}
