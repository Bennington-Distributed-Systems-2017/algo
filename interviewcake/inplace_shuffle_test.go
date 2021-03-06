// Problem:
// Given a list of integers, shuffle its location in-place.
//
// Example:
// Given:  []int{1, 2, 3, 4, 5}
// Return: []int{2, 1, 4, 3, 5}
//
// Solution:
// Iterate through the list, swap current value with a value in a randomized
// index that is between the current and last index.
//
// Cost:
// O(n) time, O(1) space.

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestInplaceShuffle(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
	}

	for _, tt := range tests {
		result := inplaceShuffle(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func inplaceShuffle(list []int) []int {
	// check edge case.
	if len(list) <= 1 {
		return list
	}

	lastIndex := len(list) - 1
	for i := 0; i < len(list); i++ {
		// get a andomized index that is between the current and last index.
		randomIndex := common.Random(i, lastIndex)

		// swap current value.
		if i != randomIndex {
			tmp := list[i]
			list[i] = list[randomIndex]
			list[randomIndex] = tmp
		}
	}

	return list
}
