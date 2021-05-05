package utils

import (
	"fmt"
	"reflect"
)

var errFormatString = "required field %s is empty"

func ValdiateFields(v interface{}) error {
	structMeta := reflect.ValueOf(v)
	if structMeta.Kind() == reflect.Ptr {
		structMeta = structMeta.Elem()
	}

	for i := 0; i < structMeta.NumField(); i++ {
		field := structMeta.Field(i)
		fieldMeta := structMeta.Type().Field(i)

		if validationCmd := fieldMeta.Tag.Get("validate"); validationCmd != "required" {
			continue
		}

		if field.Kind() == reflect.Ptr && field.IsNil() {
			return fmt.Errorf(errFormatString, fieldMeta.Name)
		}
	}
	return nil
}
