package validators

import "reflect"

type NotNilValidator struct {
	value reflect.Value
}

// GetErrorMessage implements IValidator
func (*NotNilValidator) GetErrorMessage() string {
	return "The field cannot be nil"
}

// Validate implements IValidator
func (n *NotNilValidator) Validate() bool {
	return !n.value.IsNil()
}

func NewNotNilValidator(value reflect.Value) IValidator {
	return &NotNilValidator{
		value: value,
	}
}
