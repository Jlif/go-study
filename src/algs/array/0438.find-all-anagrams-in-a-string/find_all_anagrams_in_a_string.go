package _438_find_all_anagrams_in_a_string

/**
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）
*/

func findAnagrams(s string, p string) []int {
	set := make(map[byte]int)
	window := make(map[byte]int)
	for _, v := range []byte(p) {
		set[v]++
	}

	left, right, valid := 0, 0, 0

	var res []int

	for right < len(s) {
		rs := s[right]
		if _, exist := set[rs]; exist {
			window[rs]++
			if window[rs] == set[rs] {
				valid++
			}
		}
		right++

		for right-left >= len(p) {
			if valid == len(set) {
				res = append(res, left)
			}
			ls := s[left]
			if _, exist := set[ls]; exist {
				if window[ls] == set[ls] {
					valid--
				}
				window[ls]--
			}
			left++
		}
	}

	return res
}
