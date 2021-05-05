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
		if field.Kind() == reflect.Ptr && field.IsNil() {
			return fmt.Errorf(errFormatString, structMeta.Type().Field(i).Name)
		}
	}
	return nil
}
