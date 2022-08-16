package validators

type BaseValidator struct {
	errorMessage string
}

func (b BaseValidator) GetErrorMessage() string {
	return b.errorMessage
}
