package conv

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var whitespaceRegexp = regexp.MustCompile(`[ \t\n\r]+`)
var bulletRegexp = regexp.MustCompile(`[0-9\.\*]*`)

// SquishString returns the string
// first removing all whitespace on both ends of the string,
// and then changing remaining consecutive whitespace groups into one space each.
func SquishString[T ~string](s T) T {
	str := strings.TrimSpace(string(s))
	str = whitespaceRegexp.ReplaceAllString(str, " ")
	return T(str)
}

func SquishStringFields(i any) {
	squishStringFields(reflect.ValueOf(i))
}

func squishStringFields(v reflect.Value) {
	v = IndirectReadableValue(v)
	switch v.Kind() {
	case reflect.Struct:
		squishStructStringFields(v)
	case reflect.String:
		fmt.Println(v.CanSet(), v.String())
		if v.CanSet() {
			v.SetString(SquishString(v.String()))
		}
	default:
		break
	}
}

func squishStructStringFields(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		if !fv.IsValid() || !fv.CanSet() {
			continue
		}
		switch fv.Kind() {
		case reflect.String:
			fv.SetString(SquishString(fv.String()))
		case reflect.Struct:
			squishStructStringFields(fv)
		case reflect.Ptr, reflect.Interface:
			if fv.IsNil() {
				break
			}
			squishStringFields(fv.Elem())
		default:
			break
		}
	}
}

func RemoveAllSpaces[T ~string](s T) T {
	return T(strings.ReplaceAll(string(SquishString(s)), " ", ""))
}

func RemoveBullet[T ~string](s T) T {
	s = SquishString(s)
	a := strings.Split(string(s), " ")

	if len(a) == 0 {
		return ""
	}
	a[0] = bulletRegexp.ReplaceAllString(a[0], "")
	if a[0] == "" {
		a = a[1:]
	}
	return T(strings.Join(a, " "))
}
