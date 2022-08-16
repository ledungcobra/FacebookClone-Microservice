package validators

import (
	"fmt"
)

type MinValidator struct {
	min  int
	number int
}

func (m MinValidator) GetErrorMessage() string {
	return fmt.Sprintf("Min value of the field is %d", m.min)
}

func (m MinValidator) Validate() bool {
	return m.min <= m.number
}

func NewMinValidator(number int, min int) IValidator {
	return &MinValidator{
		min:  min,
		number: number,
	}
}
