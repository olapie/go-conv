package conv_test

import (
	"testing"

	"code.olapie.com/conv"
)

func TestReverseArray(t *testing.T) {
	t.Run("IntArray", func(t *testing.T) {
		a := []int{1, 2, 3, -9, 10, 1, 101}
		conv.ReverseArray(a)
		equal(t, []int{101, 1, 10, -9, 3, 2, 1}, a)

		a = []int{1}
		conv.ReverseArray(a)
		equal(t, []int{1}, a)

		a = []int{}
		conv.ReverseArray(a)
		equal(t, []int{}, a)

		a = []int{1, 3}
		conv.ReverseArray(a)
		equal(t, []int{3, 1}, a)
	})

	t.Run("StringArray", func(t *testing.T) {
		a := []string{"a", "b", "c", "d"}
		conv.ReverseArray(a)
		equal(t, []string{"d", "c", "b", "a"}, a)
	})
}
