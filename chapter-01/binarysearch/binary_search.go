package binarysearch

import (
	"golang.org/x/exp/constraints"
)

func SearchIndex[T constraints.Ordered](array []T, element T) int {
	startIndex := 0
	endIndex := len(array) - 1

	for startIndex <= endIndex {
		middleIndex := (startIndex + endIndex) / 2
		guess := array[middleIndex]

		if guess == element {
			return middleIndex
		}

		if guess > element {
			endIndex = middleIndex - 1
		} else {
			startIndex = middleIndex + 1
		}
	}
	return -1
}
