package validators

import (
	"fmt"
	"log"
)

type LengthValidator struct {
	min   int
	max   int
	value string
}

// GetErrorMessage implements IValidator
func (l *LengthValidator) GetErrorMessage() string {
	return fmt.Sprintf("String must have the length between %d and %d", l.min, l.max)
}

// Validate implements IValidator
func (l *LengthValidator) Validate() bool {
	length := len(l.value)
	return l.min <= length && length <= l.max
}

func NewLengthValidator(value string, min, max int) IValidator {
	if min > max {
		log.Panic("Min cannot greater than max")
	}
	return &LengthValidator{
		min:   min,
		max:   max,
		value: value,
	}
}
