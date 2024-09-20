package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assertSliceEqual(t, got, want)
}

func TestSumAllTail(t *testing.T) {
	t.Run("make the sums of the tails of", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{5, 9}

		assertSliceEqual(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{1, 2})
		want := []int{0, 2}

		assertSliceEqual(t, got, want)
	})
}

func assertSliceEqual(t testing.TB, got, want []int) {
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
