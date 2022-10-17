package conv_test

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"code.olapie.com/conv"
)

func TestToInt64(t *testing.T) {
	type Bool bool
	type Int int
	goodCases := []struct {
		Value  any
		Result int64
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
			int8(123),
			123,
		}, {
			int16(123),
			123,
		}, {
			int32(123),
			123,
		}, {
			int64(123),
			123,
		}, {
			uint8(123),
			123,
		}, {
			uint16(123),
			123,
		}, {
			uint32(123),
			123,
		}, {
			uint64(123),
			123,
		}, {
			math.MaxInt64,
			math.MaxInt64,
		}, {
			math.MinInt64,
			math.MinInt64,
		}, {
			0.0,
			0,
		},
		{
			-123.1,
			-123,
		},
		{
			Bool(false),
			0,
		},
		{
			[]byte("12"),
			12,
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToInt64(c.Value)
			if err != nil {
				t.Error(err, c.Value)
			}
			equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []any{
		"hello", time.Now(), Foo{}, uint64(math.MaxInt64 + 1), uint64(math.MaxUint64),
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToInt64(c)
			if err == nil {
				t.Error("should fail", c)
			}

			if res != 0 {
				t.Error("expect empty")
			}
		}
	})
}

func TestToInt(t *testing.T) {
	type Bool bool
	type Int int
	goodCases := []struct {
		Value  any
		Result int
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
			int8(123),
			123,
		}, {
			int16(123),
			123,
		}, {
			int32(123),
			123,
		}, {
			int64(123),
			123,
		}, {
			uint8(123),
			123,
		}, {
			uint16(123),
			123,
		}, {
			uint32(123),
			123,
		}, {
			uint64(123),
			123,
		}, {
			0.0,
			0,
		},
		{
			-123.1,
			-123,
		},
		{
			Bool(false),
			0,
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToInt(c.Value)
			if err != nil {
				t.Error(err, c.Value)
			}
			equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []any{
		"hello", time.Now(), Foo{}, uint64(math.MaxInt64 + 1), uint64(math.MaxUint64),
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToInt(c)
			if err == nil {
				t.Error("should fail", c)
			}
			if res != 0 {
				t.Error("expect empty")
			}
		}
	})
}

func TestToUint64(t *testing.T) {
	type Bool bool
	type Int int
	goodCases := []struct {
		Value  any
		Result uint64
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
			int8(123),
			123,
		}, {
			int16(123),
			123,
		}, {
			int32(123),
			123,
		}, {
			int64(123),
			123,
		}, {
			uint8(123),
			123,
		}, {
			uint16(123),
			123,
		}, {
			uint32(123),
			123,
		}, {
			uint64(123),
			123,
		}, {
			math.MaxInt64,
			math.MaxInt64,
		}, {
			0.0,
			0,
		},
		{
			Bool(false),
			0,
		},
		{
			[]byte("12"),
			12,
		},
	}

	t.Run("Good", func(t *testing.T) {
		for _, c := range goodCases {
			res, err := conv.ToUint64(c.Value)
			if err != nil {
				t.Error(err, c.Value)
			}
			equal(t, c.Result, res)
		}
	})

	type Foo struct{}
	badCases := []any{
		"hello", time.Now(), Foo{}, "18446744073709551616", -1,
	}
	t.Run("Bad", func(t *testing.T) {
		for _, c := range badCases {
			res, err := conv.ToUint64(c)
			if err == nil {
				t.Error("should fail", c)
			}
			if res != 0 {
				t.Error("expect empty")
			}
		}
	})
}

func TestToIntSlice(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		l := []any{"1", 12, -13.9, json.Number("100")}
		res, err := conv.ToIntSlice(l)
		if err != nil {
			t.Error(err)
		}
		equal(t, []int{1, 12, -13, 100}, res)
	})
	t.Run("nil", func(t *testing.T) {
		res, err := conv.ToIntSlice(nil)
		if err != nil {
			t.Error(err)
		}
		if len(res) != 0 {
			t.Error("expect empty")
		}
	})
}
