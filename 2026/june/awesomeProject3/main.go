package main

func mergeArr(nums1 []int, nums2 []int) []int {
	p1 := 0
	p2 := 0
	nums3 := []int{}
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			nums3 = append(nums3, nums1[p1])
			p1++
		}
		if nums1[p1] > nums2[p2] {
			nums3 = append(nums3, nums2[p2])
			p2++
		}
		if nums1[p1] == nums2[p2] {
			nums3 = append(nums3, nums1[p1])
			p1++
			p2++
		}
	}
	return nums3
}
func main() {
}
