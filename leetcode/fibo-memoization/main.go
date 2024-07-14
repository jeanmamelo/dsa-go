package main

import "fmt"

var cache = make(map[int]int)

func fibonacciCache(n int) int {
	if val, ok := cache[n]; ok {
		return val
	}

	if n < 3 {
		return n
	}

	cache[n] = fibonacciCache(n-1) + fibonacciCache(n-2)
	return cache[n]
}

func main() {
	input := 13
	output := fibonacciCache(input)
	fmt.Println(output)
}
