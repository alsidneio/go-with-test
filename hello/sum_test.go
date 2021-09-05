package main

import (
	"reflect"
	"testing"
)

func Test_Sum(t *testing.T) {
	assert := func(t testing.TB, expected int, result int, numbers []int) {
		t.Helper() // helps with the stack trace of what line the failure is actually located
		if result != expected {
			t.Errorf("returned: %d, expected: %d, given: %v", expected, result, numbers)
		}
	}
	t.Run("Summing a collection of a fixed size(Array)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expected := 15

		assert(t, expected, result, numbers)
	})
	t.Run("Summing a collection of variable size(slice)", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sum(numbers)
		expected := 6

		assert(t, expected, result, numbers)
	})
}

func Test_SumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got: %v, wanted: %v", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
