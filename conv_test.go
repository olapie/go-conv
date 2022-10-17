package conv_test

import (
	"reflect"
	"testing"
	"time"

	"code.olapie.com/conv"
)

func equal(t *testing.T, expected, result any) {
	if reflect.DeepEqual(expected, result) {
		return
	}

	t.Errorf("expect: %v, got: %v", expected, result)
}

func TestIndirectKind(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		k := conv.IndirectKind(nil)
		equal(t, reflect.Invalid, k)
	})

	t.Run("Struct", func(t *testing.T) {
		var p time.Time
		k := conv.IndirectKind(p)
		equal(t, reflect.Struct, k)
	})

	t.Run("PointerToStruct", func(t *testing.T) {
		var p *time.Time
		k := conv.IndirectKind(p)
		equal(t, reflect.Struct, k)
	})

	t.Run("PointerToPointerToStruct", func(t *testing.T) {
		var p **time.Time
		k := conv.IndirectKind(p)
		equal(t, reflect.Struct, k)
	})

	t.Run("Map", func(t *testing.T) {
		var p map[string]any
		k := conv.IndirectKind(p)
		equal(t, reflect.Map, k)
	})

	t.Run("PointerToMap", func(t *testing.T) {
		var p map[string]any
		k := conv.IndirectKind(p)
		equal(t, reflect.Map, k)
	})
}
