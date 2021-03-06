// Problem:
// Given a list of integers, return the highest product of three numbers.
//
// Example:
// Given:  []int{-10, -10, 1, 3, 2}
// Return: 300, because -10.-10.3 gives the highest product
//
// Solution:
// Use a greedy approach to keep track of the current highest, current lowest,
// highest of three, highest of two and lowest of two for every value as we
// iterate through the list.
// Make sure to update these in the right order, in which the highest product
// of three must be calculated first using the highest and lowest product of
// two, then the highest product of two, lowest product of two, current highest
// and current lowest.
//
// Cost:
// O(n) time, O(1) space.

package interviewcake

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestHighestProductOfThree(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{-10, -10, 1, 3, 2}, 300},
		{[]int{1, 10, -5, 1, -100}, 5000},
	}

	for _, tt := range tests {
		result := highestProductOfThree(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func highestProductOfThree(list []int) int {
	if len(list) <= 3 {
		return 0
	}

	_, highest := common.Mimax(list[0], list[1])
	lowest, _ := common.Mimax(list[0], list[1])
	highestTwo := list[0] * list[1]
	lowestTwo := list[0] * list[1]
	highestThree := list[0] * list[1] * list[2]

	for i := 2; i < len(list); i++ {
		current := list[i]

		// make sure to update each variable in the right order.
		_, highestThree = common.Mimax(highestThree, current*highestTwo, current*lowestTwo)
		_, highestTwo = common.Mimax(highestTwo, current*highest, current*lowest)
		lowestTwo, _ = common.Mimax(lowestTwo, current*highest, current*lowest)
		_, highest = common.Mimax(highest, current)
		lowest, _ = common.Mimax(lowest, current)
	}

	return highestThree
}
