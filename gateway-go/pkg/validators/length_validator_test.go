package validators

import "testing"

func TestLengthValidator_Validate(t *testing.T) {
	tests := []struct {
		name      string
		validator IValidator
		want      bool
	}{
		{
			name:      "Test length validator should success",
			validator: NewLengthValidator("test", 2, 5),
			want:      true,
		},
		{
			name:      "Test length validator should fail",
			validator: NewLengthValidator("tes", 2, 3),
			want:      true,
		},
		{
			name:      "Test length validator should fail",
			validator: NewLengthValidator("te", 2, 3),
			want:      true,
		},
		{
			name:      "Test length validator should fail",
			validator: NewLengthValidator("test", 2, 3),
			want:      false,
		},
		{
			name:      "Test length validator should fail",
			validator: NewLengthValidator("t", 2, 3),
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.Validate(); got != tt.want {
				t.Errorf("LengthValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
