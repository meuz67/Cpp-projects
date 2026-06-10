package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
