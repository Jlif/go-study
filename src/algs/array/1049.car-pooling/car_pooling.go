package _049_car_pooling

func carPooling(trips [][]int, capacity int) bool {

	return false
}

type Diff struct {
	diff []int
}

// 输入原始数组，返回差分数组
func constructDiff(source []int) Diff {
	diff := make([]int, len(source))
	diff[0] = source[0]
	for i := 1; i < len(source); i++ {
		diff[i] = source[i] - source[i-1]
	}
	return Diff{diff: diff}
}

func (d Diff) increment(i int, j int, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d Diff) getResult() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}
