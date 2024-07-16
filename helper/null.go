package helper

import (
  "reflect"
)

// IsNull checks if the given value is null
func IsNull(value any) bool {
  v := reflect.ValueOf(value)

  if value == nil {
    return true
  }

  if IsPointer(value) && v.Kind().String() == "zero" {
    return true
  }

  if IsPointer(value) && v.IsZero() {
    return true
  }

  if IsPointer(value) && v.Elem().IsZero() {
    return true
  }

  if value == nil || !v.IsValid() {
    return true
  }

  if IsString(value) && len(v.String()) == 0 {
    return true
  }

  if IsArray(value) && v.Len() == 0 {
    return true
  }

  if v.Kind() == reflect.Struct {
    return reflect.TypeOf(value).NumField() == 0
  }

  return false
}
