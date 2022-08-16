package validators

import "testing"

func TestRangeValidator_Validate(t *testing.T) {
	tests := []struct {
		name      string
		validator IValidator
		want      bool
	}{
		{name: "RangeValidator should valid", validator: NewRangeValidator(10, 10, 12), want: true},
		{name: "RangeValidator should valid", validator: NewRangeValidator(11, 10, 12), want: true},
		{name: "RangeValidator should invalid", validator: NewRangeValidator(20, 10, 12), want: false},
		{name: "RangeValidator should valid", validator: NewRangeValidator(13, 10, 12), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.Validate(); got != tt.want {
				t.Errorf("RangeValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRangevalidator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Should panic when min is not less than or equal to max")
		}
	}()

	NewRangeValidator(10, 20, 10)
}
