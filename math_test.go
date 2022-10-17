package conv_test

import (
	"testing"

	"code.olapie.com/conv"
)

func TestMax(t *testing.T) {
	t.Run("N0", func(t *testing.T) {
		v := conv.Max[int]()
		equal(t, 0, v)
	})

	t.Run("N1", func(t *testing.T) {
		v := conv.Max(10)
		equal(t, 10, v)
	})

	t.Run("N2", func(t *testing.T) {
		v := conv.Max(-0.3, 10.9)
		equal(t, 10.9, v)
	})

	t.Run("N3", func(t *testing.T) {
		v := conv.Max(-0.3, 10.9, 3.8)
		equal(t, 10.9, v)
	})
}

func TestSortNumbers(t *testing.T) {
	t.Run("N0", func(t *testing.T) {
		conv.SortNumbers([]int{}, true)
	})

	t.Run("N1", func(t *testing.T) {
		conv.SortNumbers([]int{3}, true)
	})

	t.Run("N2", func(t *testing.T) {
		a := []int{3, 2}
		conv.SortNumbers(a, true)
		equal(t, []int{2, 3}, a)
	})

	t.Run("N3", func(t *testing.T) {
		a := []int{3, 5, 2}
		conv.SortNumbers(a, true)
		equal(t, []int{2, 3, 5}, a)
	})

	t.Run("N3Desc", func(t *testing.T) {
		a := []int{3, 5, 2}
		conv.SortNumbers(a, false)
		equal(t, []int{5, 3, 2}, a)
	})
}
