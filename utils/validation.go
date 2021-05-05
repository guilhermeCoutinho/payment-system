package utils

import (
	"fmt"
	"reflect"
)

var ErrMissingRequestParam error = fmt.Errorf("required field is empty")

func ValdiateFields(v interface{}) error {
	structMeta := reflect.ValueOf(v)
	if structMeta.Kind() == reflect.Ptr {
		structMeta = structMeta.Elem()
	}

	for i := 0; i < structMeta.NumField(); i++ {
		field := structMeta.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			return ErrMissingRequestParam
		}
	}
	return nil
}
