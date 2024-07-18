package helper

import "reflect"

// IsPointer check if given value is of type bool
func IsPointer(value any) bool {
	v := reflect.ValueOf(value)

	return v.Kind() == reflect.Pointer
}

func UnwrapPointer(value any) any {
	v := reflect.ValueOf(value)

	if value == nil {
		return nil
	}

	if !IsPointer(value) {
		return nil
	}

	if v.CanInterface() {
		switch v.Elem().Kind() {
		case reflect.String:
			return v.Elem().String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return v.Elem().Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return v.Elem().Uint()
		case reflect.Float32, reflect.Float64:
			return v.Elem().Float()
		case reflect.Bool:
			return v.Elem().Bool()
		case reflect.Slice, reflect.Array:
			if v.Elem().Len() > 0 {
				return v.Elem().Slice(0, v.Elem().Len()).Interface()
			}
			return nil
		default:
			{
			}
		}
	}

	return nil
}
