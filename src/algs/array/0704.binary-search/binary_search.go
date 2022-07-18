package _704_binary_search

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 //细节一
	for left <= right {    //考虑只有一个元素的数组
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1 //细节二
		} else if nums[mid] > target {
			right = mid - 1 //细节三
		}
	}

	return -1
}
