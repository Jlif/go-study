package _710_random_pick_with_blacklist

import (
	"math/rand"
)

/*
*
给定一个整数 n 和一个 无重复 黑名单整数数组blacklist。设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个未加入黑名单blacklist的整数。
任何在上述范围内且不在黑名单blacklist中的整数都应该有 同等的可能性 被返回。

优化你的算法，使它最小化调用语言 内置 随机函数的次数。

实现Solution类:

Solution(int n, int[] blacklist)初始化整数 n 和被加入黑名单blacklist的整数
int pick()返回一个范围为 [0, n - 1] 且不在黑名单blacklist 中的随机整数

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/random-pick-with-blacklist
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Solution struct {
	w []int
}

func Constructor(n int, blacklist []int) Solution {
	w := make([]int, n)
	for i := 0; i < n; i++ {
		w[i] = i
	}
	if len(blacklist) == 0 {
		return Solution{w}
	}

	blackMap := map[int]bool{}
	for _, val := range blacklist {
		blackMap[val] = true
	}

	left, right := 0, n-1
	for left < right {
		_, existRight := blackMap[w[right]]
		for existRight {
			right--
			_, existRight = blackMap[w[right]]
		}

		_, existLeft := blackMap[w[left]]
		for !existLeft && left < right {
			left++
			_, existLeft = blackMap[w[left]]
		}

		if left < right {
			w[left], w[right] = w[right], w[left]
			left++
			right--
		}
	}

	return Solution{w[0 : n-len(blacklist)]}
}

func (this *Solution) Pick() int {
	return this.w[rand.Intn(len(this.w))]
}
