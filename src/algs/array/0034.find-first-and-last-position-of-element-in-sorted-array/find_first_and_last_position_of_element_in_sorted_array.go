package _034_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if left == len(nums) {
		return -1
	}

	if nums[left] == target {
		return left
	} else {
		return -1
	}
}

func findLast(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if left == 0 {
		return -1
	}

	if nums[left-1] == target {
		return left - 1
	} else {
		return -1
	}
}
