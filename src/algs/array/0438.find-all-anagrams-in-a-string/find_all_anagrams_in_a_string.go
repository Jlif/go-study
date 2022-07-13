package _438_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	set := make(map[byte]int)
	window := make(map[byte]int)
	for _, v := range []byte(s) {
		set[v]++
	}

	left, right, valid := 0, 0, 0

	res := make([]int, 4)

	for right < len(s) {
		b := s[right]
		right++

		if _, exist := set[b]; exist {
			window[b]++
			if window[b] == set[b] {
				valid++
			}
		}

	}

	return nil
}
