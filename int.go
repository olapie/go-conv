package conv

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

const (
	// MaxInt represents maximum int
	MaxInt = 1<<(8*unsafe.Sizeof(int(0))-1) - 1
	// MinInt represents minimum int
	MinInt = -1 << (8*unsafe.Sizeof(int(0)) - 1)
	// MaxUint represents maximum uint
	MaxUint = 1<<(8*unsafe.Sizeof(uint(0))) - 1
)

// ToInt converts i to int
func ToInt(i any) (int, error) {
	n, err := parseInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to int64", i, i)
	}
	if n > MaxInt || n < MinInt {
		return 0, strconv.ErrRange
	}
	return int(n), nil
}

// MustInt panics if ToInt(i) failed
func MustInt(i any) int {
	v, err := ToInt(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

// ToInt8 converts i to int8
func ToInt8(i any) (int8, error) {
	n, err := parseInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to int8", i, i)
	}
	if n > math.MaxInt8 || n < math.MinInt8 {
		return 0, strconv.ErrRange
	}
	return int8(n), nil
}

// MustInt8 panics if ToInt8(i) failed
func MustInt8(i any) int8 {
	v, err := ToInt8(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

// ToInt16 converts i to int16
func ToInt16(i any) (int16, error) {
	n, err := parseInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to int16", i, i)
	}
	if n > math.MaxInt16 || n < math.MinInt16 {
		return 0, strconv.ErrRange
	}
	return int16(n), nil
}

// MustInt16 panics if ToInt16(i) failed
func MustInt16(i any) int16 {
	v, err := ToInt16(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToInt32(i any) (int32, error) {
	n, err := parseInt64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to int32", i, i)
	}
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0, strconv.ErrRange
	}
	return int32(n), nil
}

func MustInt32(i any) int32 {
	v, err := ToInt32(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToInt64(i any) (int64, error) {
	return parseInt64(i)
}

func MustInt64(i any) int64 {
	v, err := parseInt64(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint(i any) (uint, error) {
	n, err := parseUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to uint", i, i)
	}
	if n > MaxUint {
		return 0, strconv.ErrRange
	}
	return uint(n), nil
}

func MustUint(i any) uint {
	v, err := ToUint(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint8(i any) (uint8, error) {
	n, err := parseUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to uint8", i, i)
	}
	if n > math.MaxUint8 {
		return 0, strconv.ErrRange
	}
	return uint8(n), nil
}

func MustUint8(i any) uint8 {
	v, err := ToUint8(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint16(i any) (uint16, error) {
	n, err := parseUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to uint16", i, i)
	}
	if n > math.MaxUint16 {
		return 0, strconv.ErrRange
	}
	return uint16(n), nil
}

func MustUint16(i any) uint16 {
	v, err := ToUint16(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint32(i any) (uint32, error) {
	n, err := parseUint64(i)
	if err != nil {
		return 0, fmt.Errorf("cannot convert %#v of type %T to uint32", i, i)
	}
	if n > math.MaxUint32 {
		return 0, strconv.ErrRange
	}
	return uint32(n), nil
}

func MustUint32(i any) uint32 {
	v, err := ToUint32(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint64(i any) (uint64, error) {
	return parseUint64(i)
}

func MustUint64(i any) uint64 {
	v, err := parseUint64(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToIntSlice(i any) ([]int, error) {
	i = Indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]int); ok {
		return l, nil
	}
	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Slice, reflect.Array:
		num := v.Len()
		res := make([]int, num)
		var err error
		for j := 0; j < num; j++ {
			res[j], err = ToInt(v.Index(j).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert index %d: %w", j, err)
			}
		}
		return res, nil
	default:
		if k, err := ToInt(i); err == nil {
			return []int{k}, nil
		}
		return nil, fmt.Errorf("cannot convert %v to slice", v.Kind())
	}
}

func MustIntSlice(i any) []int {
	v, err := ToIntSlice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToInt64Slice(i any) ([]int64, error) {
	i = Indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]int64); ok {
		return l, nil
	}

	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Slice, reflect.Array:
		num := v.Len()
		res := make([]int64, num)
		var err error
		for j := 0; j < num; j++ {
			res[j], err = parseInt64(v.Index(j).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert element at index %d: %w", i, err)
			}
		}
		return res, nil
	default:
		if k, err := ToInt64(i); err == nil {
			return []int64{k}, nil
		}
		return nil, fmt.Errorf("cannot convert %#v of type %T to []int64", i, i)
	}
}

func MustInt64Slice(i any) []int64 {
	v, err := ToInt64Slice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUintSlice(i any) ([]uint, error) {
	i = Indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]uint); ok {
		return l, nil
	}

	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Slice, reflect.Array:
		num := v.Len()
		res := make([]uint, num)
		var err error
		for j := 0; j < num; j++ {
			res[j], err = ToUint(v.Index(j).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert element at index %d: %w", i, err)
			}
		}
		return res, nil
	default:
		if ui, err := ToUint(i); err == nil {
			return []uint{ui}, nil
		}
		return nil, fmt.Errorf("cannot convert %#v of type %T to []uint", i, i)
	}
}

func MustUintSlice(i any) []uint {
	v, err := ToUintSlice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func ToUint64Slice(i any) ([]uint64, error) {
	i = Indirect(i)
	if i == nil {
		return nil, nil
	}
	if l, ok := i.([]uint64); ok {
		return l, nil
	}

	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Slice, reflect.Array:
		num := v.Len()
		res := make([]uint64, num)
		var err error
		for j := 0; j < num; j++ {
			res[j], err = parseUint64(v.Index(j).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert element at index %d: %w", i, err)
			}
		}
		return res, nil
	default:
		if ui, err := ToUint64(i); err == nil {
			return []uint64{ui}, nil
		}
		return nil, fmt.Errorf("cannot convert %#v of type %T to []uint64", i, i)
	}
}

func MustUint64Slice(i any) []uint64 {
	v, err := ToUint64Slice(i)
	if err != nil {
		log.Panic(err)
	}
	return v
}

func parseInt64(i any) (int64, error) {
	i = Indirect(i)
	if i == nil {
		return 0, errNilValue
	}
	if b, ok := i.([]byte); ok {
		i = string(b)
	}
	v := reflect.ValueOf(i)
	if IsIntValue(v) {
		return v.Int(), nil
	}

	if IsUintValue(v) {
		n := v.Uint()
		if n > math.MaxInt64 {
			return 0, strconv.ErrRange
		}
		return int64(n), nil
	}

	if IsFloatValue(v) {
		return int64(v.Float()), nil
	}

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.String:
		n, err := strconv.ParseInt(v.String(), 0, 64)
		if err == nil {
			return n, nil
		}
		if errors.Is(err, strconv.ErrRange) {
			return 0, err
		}
		if f, fErr := strconv.ParseFloat(v.String(), 64); fErr == nil {
			return int64(f), nil
		}
		return 0, err
	default:
		return 0, strconv.ErrSyntax
	}
}

func parseUint64(i any) (uint64, error) {
	i = Indirect(i)
	if i == nil {
		return 0, errNilValue
	}
	if b, ok := i.([]byte); ok {
		i = string(b)
	}
	v := reflect.ValueOf(i)
	if IsIntValue(v) {
		n := v.Int()
		if n < 0 {
			return 0, strconv.ErrRange
		}
		return uint64(n), nil
	}

	if IsUintValue(v) {
		return v.Uint(), nil
	}

	if IsFloatValue(v) {
		f := v.Float()
		if f < 0 {
			return 0, strconv.ErrRange
		}
		return uint64(f), nil
	}

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.String:
		n, err := strconv.ParseInt(v.String(), 0, 64)
		if err == nil {
			if n < 0 {
				return 0, strconv.ErrRange
			}
			return uint64(n), nil
		}
		if errors.Is(err, strconv.ErrRange) {
			return 0, err
		}
		if f, fErr := strconv.ParseFloat(v.String(), 64); fErr == nil {
			if f < 0 {
				return 0, strconv.ErrRange
			}
			return uint64(f), nil
		}
		return 0, err
	default:
		return 0, strconv.ErrSyntax
	}
}

func ToUniqueIntSlice(a []int) []int {
	m := make(map[int]struct{}, len(a))
	l := make([]int, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		l = append(l, v)
	}
	return l
}

func ToUniqueInt64Slice(a []int64) []int64 {
	m := make(map[int64]struct{}, len(a))
	l := make([]int64, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		l = append(l, v)
	}
	return l
}

const (
	kilobyte = 1 << (10 * (1 + iota)) // 1 << (10*1)
	megabyte                          // 1 << (10*2)
	gigabyte                          // 1 << (10*3)
	terabyte                          // 1 << (10*4)
	petabyte                          // 1 << (10*5)
)

func SizeToHumanReadable(size int64) string {
	if size < kilobyte {
		return fmt.Sprintf("%d B", size)
	} else if size < megabyte {
		return fmt.Sprintf("%.2f KB", float64(size)/float64(kilobyte))
	} else if size < gigabyte {
		return fmt.Sprintf("%.2f MB", float64(size)/float64(megabyte))
	} else if size < terabyte {
		return fmt.Sprintf("%.2f GB", float64(size)/float64(gigabyte))
	} else if size < petabyte {
		return fmt.Sprintf("%.2f TB", float64(size)/float64(terabyte))
	} else {
		return fmt.Sprintf("%.2f PB", float64(size)/float64(petabyte))
	}
}
