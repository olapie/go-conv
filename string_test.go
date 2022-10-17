package conv_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"code.olapie.com/conv"
)

func TestToString(t *testing.T) {
	type String string
	goodCases := []struct {
		Value  any
		Result string
	}{
		{
			"This is a string",
			"This is a string",
		},
		{
			String("Typed string"),
			"Typed string",
		},
		{
			fmt.Errorf("error string"),
			"error string",
		},
		{
			10,
			"10",
		},
		{
			-10,
			"-10",
		},
		{
			10.2,
			"10.2",
		},
		{
			-10.2,
			"-10.2",
		},
		{
			[]byte("This is a byte slice"),
			"This is a byte slice",
		},
		{
			json.Number("10e6"),
			"10e6",
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToString(c.Value)
			if err != nil {
				t.Error(err, c.Value)
			}
			equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []any{
		Foo{}, &Foo{},
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToString(c)
			if err == nil {
				t.Error("should fail", c)
			}
			equal(t, "", res)
		}
	})
}

func TestToStringSlice(t *testing.T) {
	t.Run("SingleString", func(t *testing.T) {
		s := "123"
		l, err := conv.ToStringSlice(s)
		if err != nil {
			t.Error(err)
		}
		equal(t, []string{s}, l)
	})
	t.Run("SingleInt", func(t *testing.T) {
		s := 123
		l, err := conv.ToStringSlice(s)
		if err != nil {
			t.Error(err)
		}
		equal(t, []string{fmt.Sprint(s)}, l)
	})
	t.Run("IntSlice", func(t *testing.T) {
		s := []int{123, -1, 9}
		l, err := conv.ToStringSlice(s)
		if err != nil {
			t.Error(err)
		}
		equal(t, []string{"123", "-1", "9"}, l)
	})
	t.Run("MixSlice", func(t *testing.T) {
		s := []any{123, "hello", "0x123"}
		l, err := conv.ToStringSlice(s)
		if err != nil {
			t.Error(err)
		}
		equal(t, []string{"123", "hello", "0x123"}, l)
	})
	t.Run("MixArray", func(t *testing.T) {
		s := [3]any{123, "hello", "0x123"}
		l, err := conv.ToStringSlice(s)
		if err != nil {
			t.Error(err)
		}
		equal(t, []string{"123", "hello", "0x123"}, l)
	})
}
