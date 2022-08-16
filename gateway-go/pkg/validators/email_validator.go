package validators

import "regexp"

var regex *regexp.Regexp

func init() {
	regex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{1,4}$`)
}

func NewEmailValidator(email string) IValidator {
	return &RegexValidator{
		regex: regex,
		value: email,
		BaseValidator: BaseValidator{
			errorMessage: EmailErrMessage,
		},
	}
}


