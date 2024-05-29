package main

import "fmt"

func removeEvenNumbersFromArray(numbers []int) (oddNumbers []int) {
	for _, number := range numbers {
		if number%2 != 0 {
			oddNumbers = append(oddNumbers, number)
		}
	}

	return
}

func main() {
	fmt.Println(removeEvenNumbersFromArray([]int{1, 2, 3, 4, 5, 6}))
}
