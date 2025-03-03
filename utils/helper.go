package utils

import (
	"reflect"
	"strconv"
	"strings"
)

// Get JSON field name form struct
func GetJSONFieldName(model interface{}, structField string) string {
	modelType := reflect.TypeOf(model)

	if modelType.Kind() == reflect.Pointer {
		modelType = modelType.Elem()
	}

	if field, found := modelType.FieldByName(structField); found {
		jsonTag := field.Tag.Get("json")

		if jsonTag != "" && jsonTag != "-" {
			return strings.ToLower(string(jsonTag))
		}
	}

	return strings.ToLower(string(structField))
}

// Convert string to integer
func ConverStringToInt(s string, fallback ...int) int {
	defaultFallback := 0

	if len(fallback) > 0 {
		defaultFallback = fallback[0]
	}

	value, err := strconv.Atoi(s)

	if err != nil {
		return defaultFallback
	}

	return value
}

// Convert to string
func ToString(s interface{}, fallback ...string) string {
	defaultFallback := ""

	if len(fallback) > 0 {
		defaultFallback = fallback[0]
	}

	if reflect.TypeOf(s).Kind() == reflect.String && s != "" {
		return s.(string)
	}

	return defaultFallback
}

// Conver to slice
func ToSlice[T any](s any, isAssign bool) []T {
	// Case 1: If s is already a slice
	if reflect.TypeOf(s).Kind() == reflect.Slice {
		if slice, ok := s.([]T); ok {
			return slice
		}
	}

	// Case 2: If s is a pointer of slice
	if reflect.TypeOf(s).Kind() == reflect.Pointer && reflect.TypeOf(s).Elem().Kind() == reflect.Slice {
		v := reflect.ValueOf(s).Elem()

		if v.Kind() == reflect.Slice {
			if slice, ok := v.Interface().([]T); ok {
				return slice
			}
		}
	}

	// Case 3: If s is not a slice but should return a slice of T
	if isAssign {
		return make([]T, 0)
	}

	return nil
}
