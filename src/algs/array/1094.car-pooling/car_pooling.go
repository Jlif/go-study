package _094_car_pooling

import "tech.jiangchen/go-study/src/algs/data_structure"

/**
车上最初有capacity个空座位。车只能向一个方向行驶（也就是说，不允许掉头或改变方向）

给定整数capacity和一个数组 trips , trip[i] = [numPassengersi, fromi, toi]表示第 i 次旅行有numPassengersi乘客，接他们和放
他们的位置分别是fromi和toi。这些位置是从汽车的初始位置向东的公里数。

当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false

注意：
0 <= fromi < toi <= 1000
*/

func carPooling(trips [][]int, capacity int) bool {
	diff := data_structure.ConstructDiff(make([]int, 10))
	for _, v := range trips {
		// 注意这个减一，第 trip[2] 站乘客已经下车，
		// 即乘客在车上的区间是 [trip[1], trip[2] - 1]
		diff.Increment(v[1], v[2]-1, v[0])
	}
	result := diff.GetResult()

	for _, v := range result {
		if v > capacity {
			return false
		}
	}
	return true
}
