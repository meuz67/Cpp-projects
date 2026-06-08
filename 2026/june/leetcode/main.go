package main

func twoSum(arr []int, target int) []int {
	hash := make(map[int]int)
	for i, val := range arr {
		comp := target - val
		if idx, isFound := hash[comp]; isFound {
			return []int{idx, i}
		}
		hash[val] = i
	}
	return []int{-1, -1}
}
func main() {

}
