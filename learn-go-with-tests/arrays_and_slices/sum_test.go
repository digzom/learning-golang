package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSum2(t *testing.T) {
	t.Run("slice of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum2(numbers)
		want := 21
		assertIntSums(t, numbers, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("varying number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("return the sum of the tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 5}, []int{4, 33, 29, 1}, []int{49, 3, 0, 22}, []int{1, 2})
		want := []int{10, 63, 25, 2}

		assertSums(t, got, want)
	})

	t.Run("returns zero when dealing with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5, 3})
		want := []int{0, 8}

		assertSums(t, got, want)
	})
}

func assertSums(t *testing.T, got, want []int) {
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertIntSums(t *testing.T, numbers []int, got, want int) {
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
