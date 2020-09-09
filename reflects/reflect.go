package reflects

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	var numberOfValues int
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	}
	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
