package utils

import (
	"reflect"
	"strings"
)

func getJSONFieldName(model interface{}, structField string) string {
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