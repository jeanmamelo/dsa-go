package main

import "fmt"

/*
Entregue todas as permutações (grupos de combinações) possíveis de um array de números.

Exemplos:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3][2,3,1][3,1,2][3,2,1]]

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Input: nums = [1]
Output: [[1]]

Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
Todos os números são únicos.
*/

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{
			{nums[0]},
		}
	}

	if len(nums) == 2 {
		return [][]int{
			{nums[0], nums[1]},
			{nums[1], nums[0]},
		}
	}

	if len(nums) == 3 {
		return [][]int{
			{nums[0], nums[1], nums[2]},
			{nums[0], nums[2], nums[1]},
			{nums[1], nums[0], nums[2]},
			{nums[1], nums[2], nums[0]},
			{nums[2], nums[0], nums[1]},
			{nums[2], nums[1], nums[0]},
		}
	}

	permutations := [][]int{}
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	permutationsWithoutIndex := [][]int{}
	for i := 0; i < len(nums); i++ {
		permutationsWithoutIndex = permute(numsCopy)
		for j := 0; j < len(permutationsWithoutIndex); j++ {
			items := permutationsWithoutIndex[j]
			permutations = append(permutations, items)
		}
	}
	return permutations
}

func main() {
	input := []int{1, 2, 3, 4}
	output := permute(input)
	fmt.Println(output)
}
