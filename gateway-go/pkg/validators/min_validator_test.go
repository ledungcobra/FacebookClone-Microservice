package validators

import (
	"strings"
	"testing"
)

func TestMinValidator_Validate(t *testing.T) {
	tests := []struct {
		name      string
		validator IValidator
		want      bool
	}{
		{
			name:      "Test min validator should valid",
			validator: NewMinValidator(10, 10),
			want:      true,
		},
		{
			name:      "Test min validator should not valid",
			validator: NewMinValidator(9, 10),
			want:      false,
		},
		{
			name:      "Test min validator should valid",
			validator: NewMinValidator(11, 10),
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.Validate(); got != tt.want {
				if !strings.Contains(tt.validator.GetErrorMessage(), "Min") {
					t.Error("Error message should contain Min key word")
				}
				t.Errorf("MinValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
