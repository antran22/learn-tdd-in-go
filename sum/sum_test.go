package sum

import (
	"reflect"
	"testing"
)

func assertEqualInts(t testing.TB, expected int, got int) {
	t.Helper()
	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func assertEqualIntArrays(t testing.TB, expected []int, got []int) {
	t.Helper()
	expectedLength := len(expected)
	gotLength := len(got)
	if expectedLength == 0 && gotLength == 0 {
		return
	}

	if expectedLength != gotLength || !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSum(t *testing.T) {
	t.Run("test non empty array", func(t *testing.T) {
		inputs := []int{1, 2, 3, 4, 5}
		expected := 15
		got := Sum(inputs)
		assertEqualInts(t, expected, got)
	})

	t.Run("test non empty array 6 elements", func(t *testing.T) {
		inputs := []int{1, 2, 3, 4, 5, 6}
		expected := 21
		got := Sum(inputs)
		assertEqualInts(t, expected, got)
	})

	t.Run("test empty array", func(t *testing.T) {
		inputs := make([]int, 0)
		expected := 0
		got := Sum(inputs)
		assertEqualInts(t, expected, got)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("test single array", func(t *testing.T) {
		input1 := []int{1, 2, 3}
		expected := []int{6}
		got := SumAll(input1)
		assertEqualIntArrays(t, expected, got)
	})

	t.Run("test two arrays", func(t *testing.T) {
		input1 := []int{1, 2, 3}
		input2 := []int{4, 5}
		got := SumAll(input1, input2)
		expected := []int{6, 9}
		assertEqualIntArrays(t, expected, got)
	})
}

func TestSumTails(t *testing.T) {
	t.Run("test single array", func(t *testing.T) {
		input := []int{1, 2, 3}
		got := SumTails(input)
		expected := []int{5}
		assertEqualIntArrays(t, expected, got)
	})

	t.Run("test two arrays", func(t *testing.T) {
		input1 := []int{1, 2, 3}
		input2 := []int{7, 3, 4}
		got := SumTails(input1, input2)
		expected := []int{5, 7}
		assertEqualIntArrays(t, expected, got)
	})

	t.Run("test with empty array", func(t *testing.T) {
		input1 := []int{1, 4, 5}
		var input2 []int

		got := SumTails(input1, input2)
		expected := []int{9, 0}
		assertEqualIntArrays(t, expected, got)

	})
}

func TestGetTails(t *testing.T) {
	t.Run("test array with 3 elements", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{2, 3}
		got := Tails(input)
		assertEqualIntArrays(t, expected, got)
	})

	t.Run("test array with 1 elements", func(t *testing.T) {
		input := []int{1}
		var expected []int
		got := Tails(input)
		assertEqualIntArrays(t, expected, got)
	})

	t.Run("test array with 1 elements", func(t *testing.T) {
		var expected, input []int
		got := Tails(input)
		assertEqualIntArrays(t, expected, got)
	})
}
