package jsonFieldsValidator

import (
	"reflect"
)

// TagName struct tag name.
const TagName string = "validator"

// TagValue struct tag enum values.
type TagValue string

const (
	// Required field is required.
	Required TagValue = "required"
	// Optional field is optional
	Optional TagValue = "optional"
)

// String string representation of the TagValue.
func (t TagValue) String() string {
	return string(t)
}

// Validate validates a JSON request.
func Validate(obj any) bool {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tagValue := field.Tag.Get(TagName)

		if field.IsExported() && tagValue == Required.String() {
			switch val.Field(i).Kind().String() {
			case "struct":
				return Validate(val.Field(i).Interface())
			default:
				if val.Field(i).IsZero() {
					return false
				}
			}
		}

	}

	return true
}
