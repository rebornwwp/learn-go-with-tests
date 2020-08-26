package arrayslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	got := Sum(a)
	expected := 15

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{3, 4, 5})
		want := []int{0, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
