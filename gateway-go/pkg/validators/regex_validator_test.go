package validators

import "testing"

func TestRegexValidator_Validate(t *testing.T) {
	tests := []struct {
		name      string
		validator IValidator
		result    bool
	}{
		{
			name: "Test for string 192.168.1.1 should true",
			validator: NewRegexValidator(`^(\d{1,3}\.){3}\d{1,3}$`, "192.168.1.1",
				"The pattern is not corrent"),
			result: true,
		},
		{
			name: "Test for string 192.168.1.1a should fail",
			validator: NewRegexValidator(`^(\d{1,3}\.){3}\d{1,3}$`, "192.168.1.1a",
				"The pattern is not corrent"),
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.Validate(); got != tt.result {
				t.Errorf("RegexValidator.Validate() = %v, want %v", got, tt.result)
			}
		})
	}
}
