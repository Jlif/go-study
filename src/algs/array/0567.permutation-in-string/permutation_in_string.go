package _567_permutation_in_string

/**
给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func checkInclusion(s1 string, s2 string) bool {
	set := make(map[rune]int)
	window := make(map[rune]int)
	for _, v := range []rune(s1) {
		set[v]++
	}

	left, right, valid := 0, 0, 0
	src := []rune(s2)

	for right < len(src) {
		r := src[right]
		right++
		//进行窗口内数据的一系列更新
		if _, exist := set[r]; exist {
			window[r]++
			if window[r] == set[r] {
				valid++
			}
		}

		//判断左侧窗口是否要收缩
		for right-left >= len([]rune(s1)) {
			//在这里判断是否找到了合法子串
			if valid == len(set) {
				return true
			}
			l := src[left]
			left++
			//进行窗口内的一系列数据更新
			if _, exist := set[l]; exist {
				if window[l] == set[l] {
					valid--
				}
				window[l]--
			}
		}
	}

	return false
}
