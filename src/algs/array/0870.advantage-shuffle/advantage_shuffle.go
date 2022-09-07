package _870_advantage_shuffle

import (
	"fmt"
	"sort"
)

func advantageCount(nums1 []int, nums2 []int) []int {

	fmt.Println(nums1)
	fmt.Println(nums2)

	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))

	fmt.Println(nums1)
	fmt.Println(nums2)

	left := 0
	right := len(nums1) - 1
	var res []int

	for index, num := range nums2 {
		if num >= nums1[index] {
			res[index] = nums1[left]
			left++
		} else {
			res[index] = nums1[right]
			right--
		}
	}
	return res
}
