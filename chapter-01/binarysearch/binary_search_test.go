package binarysearch

import "testing"

func TestSearchIndex(t *testing.T) {
	t.Run("should find an element index in an ordered array", func(t *testing.T) {
		got := SearchIndex([]int{-1, 0, 1, 2, 3, 4, 5}, 1)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("should get -1 when an element is not present in an array", func(t *testing.T) {
		got := SearchIndex([]int{-1, 0, 1, 2, 3, 4, 5}, 100)
		want := -1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
