package conv_test

import (
	"testing"
	"time"

	"code.olapie.com/conv"
)

func TestToBool(t *testing.T) {
	type Bool bool
	type Int int
	goodCases := []struct {
		Value  any
		Result bool
	}{
		{
			"true",
			true,
		},
		{
			"false",
			false,
		},
		{
			"TRUE",
			true,
		},
		{
			"FALSE",
			false,
		},
		{
			true,
			true,
		},
		{
			false,
			false,
		},
		{
			false,
			false,
		},
		{
			1,
			true,
		},
		{
			0,
			false,
		},
		{
			-1,
			true,
		},
		{
			10.01,
			true,
		},
		{
			-10.90,
			true,
		},
		{
			0.0,
			false,
		},
		{
			Bool(true),
			true,
		},
		{
			Bool(false),
			false,
		},
		{
			Int(100),
			true,
		},
		{
			Int(0),
			false,
		},
		{
			"t",
			true,
		},
		{
			"f",
			false,
		},
		{
			"T",
			true,
		},
		{
			"F",
			false,
		},
		{
			[]byte("true"),
			true,
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToBool(c.Value)
			if err != nil {
				t.Error(err, c.Value)
			}
			equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []any{
		"hello", time.Now(), Foo{},
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToBool(c)
			if err == nil {
				t.Error("should fail", c)
			}
			equal(t, false, res)
		}
	})
}
