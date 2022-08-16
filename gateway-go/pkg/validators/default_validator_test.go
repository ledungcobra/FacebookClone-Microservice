package validators

import "testing"

func TestDefaultValidator_Validate(t *testing.T) {
	tests := []struct {
		name string
		d    *DefaultValidator
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Validate(); got != tt.want {
				t.Errorf("DefaultValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
