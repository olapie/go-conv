package conv

import "fmt"

func MustSlice[A any, B any](a []A, f func(A) B) []B {
	b := make([]B, len(a))
	for i := range a {
		b[i] = f(a[i])
	}
	return b
}

func Slice[A any, B any](a []A, f func(A) (B, error)) ([]B, error) {
	b := make([]B, len(a))
	var err error
	for i := range a {
		b[i], err = f(a[i])
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", i, err)
		}
	}
	return b, nil
}
