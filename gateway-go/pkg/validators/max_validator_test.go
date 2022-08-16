package validators

import (
	"strings"
	"testing"
)

func TestMaxValidator_Validate(t *testing.T) {
	tests := []struct {
		name      string
		validator IValidator
		want      bool
	}{
		{
			name:      "Test max validator should valid",
			validator: NewMaxValidator(10, 10),
			want:      true,
		},
		{
			name:      "Test min validator should  valid",
			validator: NewMaxValidator(9, 10),
			want:      true,
		},
		{
			name:      "Test min validator should not valid",
			validator: NewMaxValidator(11, 10),
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.Validate(); got != tt.want {
				if !strings.Contains(tt.validator.GetErrorMessage(), "Max") {
					t.Error("Error message should contain `Max` keyword")
				}
				t.Errorf("MaxValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
