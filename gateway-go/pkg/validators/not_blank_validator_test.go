package validators

import "testing"

func TestNotBlankValidator_Validate(t *testing.T) {
	tests := []struct {
		name      string
		validator IValidator
		want      bool
	}{
		{name: "Test not blank validator should success", validator: NewNotBlankValidator("test"), want: true},
		{name: "Test not blank validator should fail", validator: NewNotBlankValidator(""), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.Validate(); got != tt.want {
				t.Errorf("NotBlankValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
