package conv

import (
	"errors"
	"fmt"
	"reflect"
)

var errNilValue = errors.New("value is nil")

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// Indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func Indirect(a any) any {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// IndirectToStringerOrError returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil) or an implementation of fmt.Stringer
// or error,
func IndirectToStringerOrError(a any) any {
	if a == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func IndirectReadableValue(v reflect.Value) reflect.Value {
	for (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && !v.IsNil() {
		v = v.Elem()
	}
	return v
}

func IndirectWritableValue(v reflect.Value, populate bool) reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			if populate && v.CanSet() {
				v.Set(reflect.New(v.Type().Elem()))
			} else {
				break
			}
		}
		v = v.Elem()
	}
	if !v.CanSet() {
		panic(fmt.Sprintf("Cannot set %v", v.Kind()))
	}
	return v
}

func IndirectKind(i any) reflect.Kind {
	switch v := i.(type) {
	case reflect.Type:
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		return v.Kind()
	case reflect.Value:
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		return v.Kind()
	case reflect.Kind:
		return v
	case nil:
		return reflect.Invalid
	default:
		return IndirectKind(reflect.TypeOf(i))
	}
}

func VarargsToSlice(keyValues ...any) (keys []string, values []any, err error) {
	n := len(keyValues)
	if n%2 != 0 {
		err = errors.New("keyValues should be pairs of (string, any)")
		return
	}

	keys, values = make([]string, 0, n/2), make([]any, 0, n/2)
	for i := 0; i < n/2; i++ {
		if k, ok := keyValues[2*i].(string); !ok {
			err = fmt.Errorf("keyValues[%d] isn't convertible to string", i)
			return
		} else if keyValues[2*i+1] == nil {
			err = fmt.Errorf("keyValues[%d] is nil", 2*i+1)
			return
		} else {
			keys = append(keys, k)
			values = append(values, keyValues[2*i+1])
		}
	}
	return
}

func wrapError(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	s := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s: %w", s, err)
}

type NumberOrString interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

func IsString(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.String
}

func IsBool(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Bool
}

func IsFloat(v any) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func IsUint(v any) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func IsInt(v any) bool {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func IsNumber(v any) bool {
	return IsInt(v) || IsUint(v) || IsFloat(v)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
