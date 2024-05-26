package main

import (
	"fmt"
	"math"
)

// Write a function that takes a non-empty array of integers and returns the maximum sum
// that can be obtained by summing up all of the integers in a non-empty subarray of the input array.
// A subarray must only contain adjacent numbers.

// Input: [3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4]
// Output: 19

// Input: [-2, -1]
// Output: -1

// Input: [1, 2, -4, 3, 5, -9, 8, 1, 2]
// Output: 11

// Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// Output: 6

// Input: [1]
// Output: 1

// Input: [5, 4, -1, 7, 8]
// Output: 23

func maxSubArray(nums []int) int {
	var maxSum int = math.MinInt32
	var currentSum int = 0
	for _, num := range nums {
		currentSum = max(num, num+currentSum)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}
	// input := []int{-2, -1}
	// input := []int{1, 2, -4, 3, 5, -9, 8, 1, 2}
	// input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// input := []int{1}
	// input := []int{5, 4, -1, 7, 8}
	output := maxSubArray(input)
	fmt.Println(output)
}
