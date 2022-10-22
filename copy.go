package conv

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func GobCopy(dst, src any) error {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(src)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	dec := gob.NewDecoder(&b)
	err = dec.Decode(dst)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}
	return nil
}

func JSONCopy(dst, src any) error {
	b, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	err = json.Unmarshal(b, dst)
	if err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	return nil
}

func DeepCopy(dst, src any) error {
	if dst == nil {
		return errors.New("dst cannot be nil")
	}

	if src == nil {
		return errors.New("src cannot be nil")
	}

	dstType := reflect.TypeOf(dst)
	srcType := reflect.TypeOf(src)

	dstKind := IndirectReadableValue(reflect.ValueOf(dst)).Kind()
	srcKind := IndirectReadableValue(reflect.ValueOf(src)).Kind()
	if reflect.PtrTo(srcType).ConvertibleTo(dstType) || (srcKind == dstKind && dstKind == reflect.Struct) {
		err := GobCopy(dst, src)
		if err != nil {
			return fmt.Errorf("json copy: %w", err)
		}
		return nil
	}

	return fmt.Errorf("cannot copy %T to %T", src, dst)
}

func DeepNew(t reflect.Type) reflect.Value {
	v := reflect.New(t)
	e := v.Elem()
	for e.Kind() == reflect.Ptr {
		e.Set(reflect.New(e.Type().Elem()))
		e = e.Elem()
	}

	if e.Kind() != reflect.Struct {
		return v
	}

	for i := 0; i < e.NumField(); i++ {
		ft := e.Type().Field(i)
		if !e.Field(i).CanSet() {
			continue
		}

		switch ft.Type.Kind() {
		case reflect.Ptr:
			e.Field(i).Set(DeepNew(ft.Type.Elem()))
		case reflect.Struct:
			e.Field(i).Set(DeepNew(ft.Type).Elem())
		case reflect.Slice:
			elemVal := DeepNew(ft.Type.Elem()).Elem()
			sliceVal := reflect.New(ft.Type).Elem()
			sliceVal = reflect.Append(sliceVal, elemVal)
			e.Field(i).Set(sliceVal)
		}
	}
	return v
}

func CloneSlice[T any](a []T) []T {
	res := make([]T, len(a))
	for i, v := range a {
		res[i] = v
	}
	return res
}

func CloneMap[K comparable, V any](m map[K]V) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}
