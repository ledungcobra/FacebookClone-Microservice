package validators

import "testing"

func TestInValidator_Validate(t *testing.T) {
	tests := []struct {
		name      string
		validator IValidator
		want      bool
	}{
		{name: "Test in validator should success", validator: NewInValidator("test", []string{"test", "test2"}), want: true},
		{name: "Test in validator should fail", validator: NewInValidator("tes3t", []string{"test", "test2"}), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.Validate(); got != tt.want {
				t.Errorf("InValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
