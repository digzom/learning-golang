package sum

import "testing"

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

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
