package _187_repeated_dna_sequences

/**
DNA序列由一系列核苷酸组成，缩写为'A','C','G'和'T'。

例如，"ACGAATTCCG"是一个 DNA序列 。
在研究 DNA 时，识别 DNA 中的重复序列非常有用。

给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的长度为10的序列(子字符串)。你可以按 任意顺序 返回答案。
*/

func findRepeatedDnaSequences(s string) []string {
	arr := make([]byte, len(s))
	for i, v := range s {
		switch string(v) {
		case "A":
			arr[i] = 0
		case "C":
			arr[i] = 1
		case "G":
			arr[i] = 2
		case "T":
			arr[i] = 3
		}
	}

	left, right := 0, 0
	var res []string
	resMap := map[string]struct{}{}
	seek := map[int]bool{}

	tmp := 0
	for right < len(arr) {
		tmp = addNewStr(tmp, arr[right])
		right++

		for right-left >= 10 {
			if _, exist := seek[tmp]; exist {
				resMap[s[left:right]] = struct{}{}
			} else {
				seek[tmp] = true
			}
			left++
		}
	}

	for key := range resMap {
		res = append(res, key)
	}

	return res
}

func addNewStr(num int, b byte) int {
	num = num << 2
	num = num & 0xfffff
	num = num | int(b)
	return num
}
