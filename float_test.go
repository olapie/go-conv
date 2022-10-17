package conv_test

import (
	"math"
	"testing"
	"time"

	"code.olapie.com/conv"
)

func TestToFloat64(t *testing.T) {
	type Bool bool
	type Int int
	type Float64 float64
	goodCases := []struct {
		Value  any
		Result float64
	}{
		{
			true,
			1,
		}, {
			false,
			0,
		}, {
			123456,
			123456,
		}, {
			Int(123456),
			123456,
		}, {
			int8(8),
			8,
		}, {
			int16(16),
			16,
		}, {
			int32(32),
			32,
		}, {
			int64(64),
			64,
		}, {
			uint8(80),
			80,
		}, {
			uint16(160),
			160,
		}, {
			uint32(320),
			320,
		}, {
			uint64(123),
			123,
		}, {
			math.MaxFloat64,
			math.MaxFloat64,
		}, {
			-math.MaxFloat64,
			-math.MaxFloat64,
		}, {
			0.0,
			0,
		},
		{
			-123.1,
			-123.1,
		},
		{
			Bool(false),
			0,
		},
		{
			[]byte("123.309"),
			123.309,
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToFloat64(c.Value)
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
			res, err := conv.ToFloat64(c)
			if err == nil {
				t.Error("should fail", c)
			}
			if res != 0 {
				t.Error("expect empty")
			}
		}
	})
}
