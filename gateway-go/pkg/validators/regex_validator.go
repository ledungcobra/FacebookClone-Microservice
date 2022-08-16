package validators

import (
	"regexp"
)

type RegexValidator struct {
	regex *regexp.Regexp
	value string
	BaseValidator
}

func NewRegexValidator(pattern, value string, errorMsg string) IValidator {
	regex := regexp.MustCompile(pattern)
	return &RegexValidator{
		regex: regex,
		value: value,
		BaseValidator: BaseValidator{
			errorMessage: errorMsg,
		},
	}
}

// Validate implements Validator
func (e *RegexValidator) Validate() bool {
	return e.regex.Match([]byte(e.value))
}
