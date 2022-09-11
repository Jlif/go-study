package _528_random_pick_with_weight

import (
	"math/rand"
	"sort"
)

/**
给你一个 下标从 0 开始 的正整数数组w ，其中w[i] 代表第 i 个下标的权重。

请你实现一个函数pickIndex，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i的 概率 为 w[i] / sum(w) 。

例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3)= 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3)= 0.75（即，75%）。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/random-pick-with-weight
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Solution struct {
	pre []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (s *Solution) PickIndex() int {
	x := rand.Intn(s.pre[len(s.pre)-1]) + 1
	return sort.SearchInts(s.pre, x)
}
