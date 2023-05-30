//go:build js

package jshelpers

import (
	"fmt"
	"reflect"
	"syscall/js"
)

var (
	jsArray = js.Global().Get("Array")
)

func asSlice(typ reflect.Type, value js.Value) (any, error) {
	l := value.Length()

	slice := reflect.MakeSlice(reflect.SliceOf(typ), l, l)
	for i := 0; i < l; i++ {
		val, err := goValue(typ, value.Index(i))
		if err != nil {
			return nil, fmt.Errorf("slice item is invalid: %w", err)
		}
		slice.Index(i).Set(reflect.ValueOf(val))
	}

	return slice.Interface(), nil
}

// AsSlice converts a JS value to a slice of T
// value must be an array of a type handled by goValue
func AsSlice[T any](value js.Value) ([]T, error) {
	var t T

	if !value.InstanceOf(jsArray) {
		return nil, fmt.Errorf("value type should be Array")
	}

	slice, err := asSlice(reflect.TypeOf(t), value)
	if err != nil {
		return nil, err
	}

	return slice.([]T), nil
}
