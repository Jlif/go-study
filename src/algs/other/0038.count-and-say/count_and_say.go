package _038_count_and_say

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		s := countAndSay(n - 1)
		var t string
		var count int
		var result string
		for i := 0; i < len(s); i++ {
			st := string(s[i])
			if i == 0 {
				t = st
				count = 1
			} else {
				if st == t {
					count++
				} else {
					result = result + strconv.Itoa(count) + t
					t = st
					count = 1
				}
			}
		}
		result = result + strconv.Itoa(count) + t
		return result
	}
}
