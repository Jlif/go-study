package _003_longest_substring_without_repeating_characters

import "unicode/utf8"

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度
*/

func lengthOfLongestSubstring(s string) int {
	dict := map[rune]int{}
	n := utf8.RuneCountInString(s)
	rs := []rune(s)
	rp, result := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			//左边指针右移，移除一个字符
			delete(dict, rs[i-1])
		}
		for rp+1 < n && dict[rs[rp+1]] == 0 {
			dict[rs[rp+1]]++
			rp++
		}

		result = max(result, rp-i+1)
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
