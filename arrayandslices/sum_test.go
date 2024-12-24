package arrayandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)

	})

}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		num1 := []int{1, 2, 3, 4, 5, 6}
		num2 := []int{7, 8, 9}

		got := SumAllTails(num1, num2)
		want := []int{20, 17}

		checkSums(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int, numbers []int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	}
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	//DeepEqual is not type safe. checks if both are equal or not.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
