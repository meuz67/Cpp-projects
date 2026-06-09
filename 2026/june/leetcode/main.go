package main

import (
	"fmt"
)

func RomanToInt(s string) int {
	dictionary := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	n := len(s)
	for i := 0; i < n; i++ {
		currentVal := dictionary[s[i]]
		if i < n-1 && currentVal < dictionary[s[i+1]] {
			result -= currentVal
		} else {
			result += currentVal
		}
	}
	return result
}
func main() {
	fmt.Println(RomanToInt("XVI"))
}
