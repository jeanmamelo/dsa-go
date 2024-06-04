package main

import "fmt"

func moveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[j] != 0 {
			j++
		}
	}
	return nums
}

func main() {
	input := []int{1, 0, 3, 2, 0}
	// input := []int{0, 1, 0, 3, 12}
	// input := []int{0}
	output := moveZeroes(input)
	fmt.Println(output)
}
