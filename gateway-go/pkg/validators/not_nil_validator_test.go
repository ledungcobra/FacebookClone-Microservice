package validators

import (
	"reflect"
	"testing"
)

func TestNotNilValidator_Validate(t *testing.T) {
	type Test1 struct {
		value string
	}

	var t1 *Test1 = nil
	
	tests := []struct {
		name      string
		validator IValidator
		want      bool
	}{
		{name: "Test not nil validator should success", validator: NewNotNilValidator(reflect.ValueOf(&Test1{"hi"})), want: true},
		{name: "Test not nil validator should fail", validator: NewNotNilValidator(reflect.ValueOf(t1)), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.validator.Validate(); got != tt.want {
				t.Errorf("NotNilValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
