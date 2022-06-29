package data_structure

type Diff struct {
	diff []int
}

// ConstructDiff 输入原始数组，返回差分数组
func ConstructDiff(source []int) Diff {
	diff := make([]int, len(source))
	diff[0] = source[0]
	for i := 1; i < len(source); i++ {
		diff[i] = source[i] - source[i-1]
	}
	return Diff{diff: diff}
}

func (d Diff) Increment(i int, j int, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d Diff) GetResult() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}
