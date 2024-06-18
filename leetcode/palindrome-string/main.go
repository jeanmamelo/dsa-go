package main

import "fmt"

func isPalindrome(word string) bool {
	start := 0
	end := len(word) - 1

	for start < end {
		letterStart := fmt.Sprintf("%c", word[start])
		letterEnd := fmt.Sprintf("%c", word[end])

		if letterStart != letterEnd {
			return false
		}

		start++
		end--
	}
	return true
}

func main() {
	println(isPalindrome("natan"))
	println(isPalindrome("jean"))
}
