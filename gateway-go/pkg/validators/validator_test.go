package validators

import (
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	type args struct {
		testObject any
	}

	type User struct {
		Email string `validator:"email"`
		Regex string `validator:"regex:'^(\\d{1,3}\\.){3}\\d{1,3}$'"`
	}

	tests := []struct {
		name             string
		args             args
		isValid          bool
		validatorResults []ValidationResult
	}{
		{
			name: "Test valid",
			args: args{
				testObject: User{
					Email: "ledungcobra@gmail.com",
					Regex: "192.168.1.1",
				},
			},
			isValid: true,
		},
		{
			name: "Test regex invalid",
			args: args{
				testObject: User{
					Email: "ledungcobra@gmail.com",
					Regex: "192.168.1.1a",
				},
			},
			isValid: false,
			validatorResults: []ValidationResult{
				{fieldName: "Regex", errors: []string{RegexErrMessage}},
			},
		},
		{
			name: "Test email invalid",
			args: args{
				testObject: User{
					Email: "@gmail.com",
					Regex: "192.168.1.1",
				},
			},
			isValid: false,
			validatorResults: []ValidationResult{
				{fieldName: "Email", errors: []string{EmailErrMessage}},
			},
		},
		{
			name: "Test email and regex should invalid",
			args: args{
				testObject: User{
					Email: "@gmail.com",
					Regex: "192.168.1.1a",
				},
			},
			isValid: false,
			validatorResults: []ValidationResult{
				{fieldName: "Email", errors: []string{EmailErrMessage}},
				{fieldName: "Regex", errors: []string{RegexErrMessage}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, validatorResults := Validate(tt.args.testObject)
			if isValid != tt.isValid {
				t.Errorf("Validate() got = %v, want %v", isValid, tt.isValid)
			}
			if !reflect.DeepEqual(validatorResults, tt.validatorResults) {
				t.Errorf("Validate() got1 = %v, want %v", validatorResults, tt.validatorResults)
			}
		})
	}
}
