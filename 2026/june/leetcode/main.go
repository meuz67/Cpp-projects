package main

import (
	"fmt"
	"strconv"
)

func IsPalindrome(num int) bool {
	str := strconv.Itoa(num)
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	number := 101
	fmt.Println(IsPalindrome(number))
}
