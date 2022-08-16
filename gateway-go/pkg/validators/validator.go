package validators

import (
	"reflect"
	"strings"
)

type IValidator interface {
	Validate() bool
	GetErrorMessage() string
}

const (
	EmailErrMessage   = "Email is not valid"
	RegexErrMessage   = "Input is not valid"
	Regex             = "regex"
	Email             = "email"
	ValidtorTag       = "validator"
	TagPropSeperator  = ";"
	TagFieldSeperator = ":"
)

type ValidationResult struct {
	fieldName string
	errors    []string
}

// Validate the object passed in function by looking validator tag
// validator tag support email using regex
func Validate(object any) (bool, []ValidationResult) {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	numFields := objType.NumField()
	var validationResults []ValidationResult
	for fieldIndex := 0; fieldIndex < numFields; fieldIndex++ {
		curField, curValue := extractFieldAndValue(objType, fieldIndex, objValue)
		tagValidators := curField.Tag.Get(ValidtorTag)
		tagFieldItems := strings.Split(tagValidators, TagPropSeperator)
		validators := extractValidators(tagFieldItems, curValue)
		fieldValidatorResult := getFieldValidatorResult(validators, curField)
		if fieldValidatorResult != nil {
			validationResults = append(validationResults, *fieldValidatorResult)
		}
	}
	return len(validationResults) == 0, validationResults
}

func extractFieldAndValue(objType reflect.Type, i int, objValue reflect.Value) (reflect.StructField, reflect.Value) {
	currentField := objType.Field(i)
	curValue := objValue.Field(i)
	return currentField, curValue
}

func extractValidators(tagFieldItems []string, curValue reflect.Value) []IValidator {
	var validators []IValidator
	for _, tagFieldItem := range tagFieldItems {
		tagFieldName, tagFieldValue := extractTagFieldAndValue(strings.Split(tagFieldItem, TagFieldSeperator))
		validators = append(validators, mapToCorrespondingValidator(tagFieldName, curValue, tagFieldValue))
	}
	return validators
}

func getFieldValidatorResult(validators []IValidator, curField reflect.StructField) *ValidationResult {
	var errors []string
	for _, validator := range validators {
		if !validator.Validate() {
			errors = append(errors, validator.GetErrorMessage())
		}
	}
	if len(errors) > 0 {
		return &ValidationResult{
			fieldName: curField.Name,
			errors:    errors,
		}
	}
	return nil
}

func mapToCorrespondingValidator(tagFieldName string, curValue reflect.Value, tagFieldValue string) IValidator {
	var curValidator IValidator
	switch tagFieldName {
	case Email:
		curValidator = NewEmailValidator(curValue.String())
	case Regex:
		curValidator = NewRegexValidator(tagFieldValue, curValue.String(), RegexErrMessage)
	default:
		curValidator = defaultValidator
	}
	return curValidator
}

func extractTagFieldAndValue(tagField []string) (string, string) {
	fieldName := tagField[0]
	var fieldValue string
	if len(tagField) == 2 {
		fieldValue = strings.Replace(tagField[1], "'", "", -1)
	}
	return fieldName, fieldValue
}
