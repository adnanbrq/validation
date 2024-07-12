package helper

import "reflect"

// IsBool check if given value is of type bool
func IsBool(value interface{}) bool {
	v := reflect.ValueOf(value)

	return v.Kind() == reflect.Bool
}

func ParseBool(value string) bool {
	switch value {
	case "true", "1":
		return true
	default:
		return false
	}
}
