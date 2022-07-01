package reflection

import (
	"reflect"
)

func Walk(input interface{}, f func(a string)) {
	value := getUnderlyingValue(input)

	for fieldIndex := 0; fieldIndex < value.NumField(); fieldIndex++ {
		field := value.Field(fieldIndex)
		switch field.Kind() {
		case reflect.String:
			f(field.String())
		case reflect.Struct:
			Walk(field.Interface(), f)
		}
	}
}

func getUnderlyingValue(input interface{}) reflect.Value {
	value := reflect.ValueOf(input)
	if value.Kind() == reflect.Pointer {
		return value.Elem()
	}
	return value
}
