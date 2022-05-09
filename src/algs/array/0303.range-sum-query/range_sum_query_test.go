package main

import (
	"fmt"
	"testing"
)

type question303 struct {
	para303
	ans303
}

type para303 struct {
	nums     []int
	sumRange [2]int
}

type ans303 struct {
	ans int
}

func Test(t *testing.T) {
	qs := []question303{
		{
			para303{nums: []int{-2, 0, 3, -5, 2, -1}, sumRange: [2]int{2, 4}},
			ans303{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 303------------------------\n")

	for _, q := range qs {
		p, a := q.para303, q.ans303
		obj := Constructor(p.nums)
		if ans := obj.SumRange(p.sumRange[0], p.sumRange[1]); ans != a.ans {
			t.Errorf("SumRange(%d,%d) expected be %d, but %d got", p.sumRange[0], p.sumRange[1], a.ans, ans)
		} else {
			fmt.Printf("【input】:%v\t【output】:%d\n", p, obj.SumRange(p.sumRange[0], p.sumRange[1]))
		}
	}
	fmt.Printf("\n\n\n")
}
