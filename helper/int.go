package helper

import (
	"reflect"
	"strconv"
)

// IsInt checks if given value is Int, Int8, Int16, Int32 or Int64
func IsInt(value interface{}) bool {
	v := reflect.ValueOf(value)

	if !v.IsValid() {
		return false
	}

	kind := v.Kind()
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}

func IsUint(value interface{}) bool {
	v := reflect.ValueOf(value)

	if !v.IsValid() {
		return false
	}

	kind := v.Kind()
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64
}

func IsFloat(value interface{}) bool {
	v := reflect.ValueOf(value)

	if !v.IsValid() {
		return false
	}

	kind := v.Kind()
	return kind == reflect.Float32 || kind == reflect.Float64
}

func ParseFloat(v string) float64 {
	res, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0
	}

	return res
}

func ParseInt(v string) int64 {
	res, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		return 0
	}

	return res
}

func ParseUint(v string) uint64 {
	res, err := strconv.ParseUint(v, 0, 64)
	if err != nil {
		return 0
	}

	return res
}
