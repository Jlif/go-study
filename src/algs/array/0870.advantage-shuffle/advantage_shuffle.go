package _870_advantage_shuffle

import (
	"sort"
)

/**
给定两个大小相等的数组nums1和nums2，nums1相对于 nums的优势可以用满足nums1[i] > nums2[i]的索引 i的数目来描述。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/advantage-shuffle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)

	tmp := make(itemSlice, n)
	for index, val := range nums2 {
		tmp[index] = item{index: index, value: val}
	}

	//s1升序
	sort.Ints(nums1)
	//s2降序
	sort.Sort(tmp)

	left := 0
	right := n - 1
	res := make([]int, n)
	for _, num := range tmp {
		if num.value >= nums1[right] {
			res[num.index] = nums1[left]
			left++
		} else {
			res[num.index] = nums1[right]
			right--
		}
	}
	return res
}

type item struct {
	index int
	value int
}

type itemSlice []item

func (t itemSlice) Len() int {
	return len(t)
}
func (t itemSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t itemSlice) Less(i, j int) bool {
	return t[j].value < t[i].value
}
