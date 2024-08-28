package validation

import (
	"reflect"
)

// Function to get empty fields, non-recursive.
// Does not view the fields of nested structures
func GetEmptyFields(v any) []string {
	return getEmptyFieldsInternal(reflect.ValueOf(v), false)
}

// Function to get empty fields.
// View through all fields of nested structures
func GetEmptyFieldsRecursive(v any) []string {
	return getEmptyFieldsInternal(reflect.ValueOf(v), true)
}

// Function to get empty fields in a struct
func getEmptyFieldsInternal(objValue reflect.Value, recursive bool) []string {
	// Dereference pointer if necessary
	objValue = dereferenceValue(objValue)

	if objValue.Kind() != reflect.Struct {
		return nil
	}

	zeroFields := make([]string, 0)
	objType := objValue.Type()

	for idx := 0; idx < objValue.NumField(); idx++ {
		field := objValue.Field(idx)

		// Check if the field value can be retrieved
		if !field.CanInterface() {
			continue
		}

		// Dereference pointer or interface
		field = dereferenceValue(field)

		// Recursive check if the field is a nested struct
		if recursive && field.Kind() == reflect.Struct {
			fields := getEmptyFieldsInternal(field, true)
			zeroFields = append(zeroFields, fields...)
			continue
		}

		// Check if the field value is empty
		if isEmptyValue(field) {
			zeroFields = append(zeroFields, objType.Field(idx).Name)
		}
	}

	return zeroFields
}

// Common function to handle pointer and interface dereferencing
func dereferenceValue(val reflect.Value) reflect.Value {
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Pointer {
		if !val.IsNil() {
			val = val.Elem()
		}
	}
	return val
}

// Common function to check if the value is empty
func isEmptyValue(val reflect.Value) bool {
	return !val.IsValid() || val.IsZero()
}
