package validators

import (
	"testing"
)

func TestValidateEmailValidator(t *testing.T) {
	tests := []struct {
		name  string
		input IValidator
		want  bool
	}{
		{"Test valid email", NewEmailValidator("test@gmail.com"), true},
		{"Test invalid email", NewEmailValidator("@gmail.com"), false},
		{"Test email with multiple prefixes", NewEmailValidator("test@gmail.com.vn"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := tt.input.Validate(); result != tt.want {
				t.Errorf("Want %v but got %v", tt.want, result)
				if tt.input.GetErrorMessage() != EmailErrMessage {
					t.Error("Error message must not be empty in case of error occur")
				}
			}
		})
	}
}
