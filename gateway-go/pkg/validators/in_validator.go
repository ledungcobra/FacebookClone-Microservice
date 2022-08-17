package validators

import "fmt"

type InValidator struct {
	value    string
	patterns []string
}

func (i *InValidator) GetErrorMessage() string {
	return fmt.Sprintf("%s not in %+v", i.value, i.patterns)
}

func (i *InValidator) Validate() bool {
	for _, pattern := range i.patterns {
		if i.value == pattern {
			return true
		}
	}
	return false
}

func NewInValidator(value string, patterns []string) IValidator {
	return &InValidator{
		value:    value,
		patterns: patterns,
	}
}
