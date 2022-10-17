package conv

import (
	"fmt"
	"reflect"
)

type Validator interface {
	Validate() error
}

func Validate(i any) error {
	if v, ok := i.(Validator); ok {
		return v.Validate()
	}

	v := reflect.ValueOf(i)
	if v.IsValid() && (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && !v.IsNil() {
		v = v.Elem()
		if v.CanInterface() {
			if va, ok := v.Interface().(Validator); ok {
				return va.Validate()
			}
		}
	}

	v = IndirectReadableValue(v)
	if v.Kind() == reflect.Struct {
		t := v.Type()
		for j := 0; j < v.NumField(); j++ {
			if !isExported(t.Field(j).Name) {
				continue
			}
			if err := Validate(v.Field(j).Interface()); err != nil {
				return fmt.Errorf("%s:%w", t.Field(j).Name, err)
			}
		}
	}
	return nil
}
