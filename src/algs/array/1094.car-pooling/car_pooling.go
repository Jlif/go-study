package _094_car_pooling

/**
车上最初有capacity个空座位。车只能向一个方向行驶（也就是说，不允许掉头或改变方向）

给定整数capacity和一个数组 trips , trip[i] = [numPassengersi, fromi, toi]表示第 i 次旅行有numPassengersi乘客，接他们和放
他们的位置分别是fromi和toi。这些位置是从汽车的初始位置向东的公里数。

当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false

注意：
0 <= fromi < toi <= 1000
*/

func CarPooling(trips [][]int, capacity int) bool {
	diff := constructDiff(make([]int, 1001))
	for _, v := range trips {
		diff.increment(v[1], v[2], v[0])
	}
	result := diff.getResult()

	for _, v := range result {
		if v > capacity {
			return false
		}
	}
	return true
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
