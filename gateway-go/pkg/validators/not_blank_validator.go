package validators

type NotBlankValidator struct {
	value string
	BaseValidator
}

// Validate implements IValidator
func (f *NotBlankValidator) Validate() bool {
	return f.value != ""
}

func NewNotBlankValidator(value string) IValidator {
	return &NotBlankValidator{
		value: value,
		BaseValidator: BaseValidator{
			"Field is required",
		},
	}
}
