package validators

var defaultValidator IValidator = &DefaultValidator{}

type DefaultValidator struct {
}

// GetErrorMessage should always return empty string
func (*DefaultValidator) GetErrorMessage() string {
	return ""
}

// Validate should always return true
func (*DefaultValidator) Validate() bool {
	return true
}
