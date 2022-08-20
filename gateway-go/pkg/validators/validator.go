package validators

import (
	"reflect"
	"strings"

	"ledungcobra/gateway-go/pkg/common"
)

type IValidator interface {
	Validate() bool
	GetErrorMessage() string
}

const (
	EmailErrMessage   = "Email is not valid"
	RegexErrMessage   = "Input is not valid"
	ValidatorTag      = "validator"
	TagPropSeperator  = ";"
	TagFieldSeperator = ":"
	Regex             = "regex"
	Email             = "email"
	Min               = "min"
	Max               = "max"
	Range             = "range"
	NotBlank          = "not_blank"
	Length            = "length"
	NotNil            = "not_nil"
	In                = "in"
)

type ValidationResult struct {
	Field  string   `json:"field"`
	Errors []string `json:"errors"`
}

// Validate the object passed in function by looking validator tag
// validator tag support email using regex
func Validate(object any) (bool, []ValidationResult) {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	if objType.Kind() == reflect.Pointer {
		objType = objType.Elem()
		objValue = objValue.Elem()
	}
	numFields := objType.NumField()
	var validationResults []ValidationResult
	for fieldIndex := 0; fieldIndex < numFields; fieldIndex++ {
		curField, curValue := extractFieldAndValue(objType, objValue, fieldIndex)
		tagValidators := curField.Tag.Get(ValidatorTag)
		tagFieldItems := strings.Split(tagValidators, TagPropSeperator)
		validators := extractValidators(tagFieldItems, curValue)
		validationResult := validateCurField(validators, curField)
		if validationResult != nil {
			validationResults = append(validationResults, *validationResult)
		}
	}
	return len(validationResults) == 0, validationResults
}

func extractFieldAndValue(objType reflect.Type, objValue reflect.Value, i int) (reflect.StructField, reflect.Value) {
	return objType.Field(i), objValue.Field(i)
}

func extractValidators(tagFieldItems []string, curValue reflect.Value) []IValidator {
	var validators []IValidator
	for _, tagFieldItem := range tagFieldItems {
		tagFieldName, tagFieldValue := extractTagFieldAndValue(strings.Split(tagFieldItem, TagFieldSeperator))
		validators = append(validators, mapToCorrespondingValidator(tagFieldName, tagFieldValue, curValue))
	}
	return validators
}

func validateCurField(validators []IValidator, curField reflect.StructField) *ValidationResult {
	var errors []string
	for _, validator := range validators {
		if !validator.Validate() {
			errors = append(errors, validator.GetErrorMessage())
		}
	}
	if len(errors) > 0 {
		return &ValidationResult{
			Field:  common.ToSnakeCase(curField.Name),
			Errors: errors,
		}
	}
	return nil
}

func mapToCorrespondingValidator(tagFieldName string, tagFieldValue string, curValue reflect.Value) IValidator {
	switch tagFieldName {
	case Email:
		return NewEmailValidator(curValue.String())
	case Regex:
		return NewRegexValidator(tagFieldValue, curValue.String(), RegexErrMessage)
	case Min:
		return NewMinValidator(int(curValue.Int()), common.ToInt(tagFieldName))
	case Max:
		return NewMaxValidator(int(curValue.Int()), common.ToInt(tagFieldName))
	case Range:
		min, max := extractMinMax(tagFieldValue)
		return NewRangeValidator(int(curValue.Int()), min, max)
	case NotBlank:
		return NewNotBlankValidator(curValue.String())
	case Length:
		min, max := extractMinMax(tagFieldValue)
		return NewLengthValidator(curValue.String(), min, max)
	case NotNil:
		return NewNotNilValidator(curValue)
	case In:
		return NewInValidator(curValue.String(), toListTag(tagFieldValue))
	default:
		return defaultValidator
	}
}

func toListTag(str string) []string {
	return strings.Split(str, ",")
}
func extractMinMax(tagFieldValue string) (int, int) {
	rangeTokens := strings.Split(tagFieldValue, "-")
	min := common.ToInt(rangeTokens[0])
	max := common.ToInt(rangeTokens[1])
	return min, max
}

func extractTagFieldAndValue(tagField []string) (string, string) {
	fieldName := tagField[0]
	var fieldValue string
	if len(tagField) == 2 {
		fieldValue = strings.Replace(tagField[1], "'", "", -1)
	}
	return fieldName, fieldValue
}
