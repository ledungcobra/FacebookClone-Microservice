package validators

import "fmt"

type MaxValidator struct {
	max  int
	number int
}

func (m MaxValidator) GetErrorMessage() string {
	return fmt.Sprintf("Max value of the field is %d", m.max)
}

func (m MaxValidator) Validate() bool {
	return m.max >= m.number
}

func NewMaxValidator(number int, max int) IValidator {
	return &MaxValidator{
		max:  max,
		number: number,
	}
}
