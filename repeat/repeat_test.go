package repeat

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("five character", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		if expected != got {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})

	t.Run("six character", func(t *testing.T) {
		got := Repeat("b", 6)
		expected := "bbbbbb"

		if expected != got {
			t.Errorf("Expected %q, got %q", expected, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 8)
	}
}
