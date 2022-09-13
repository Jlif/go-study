package _710_random_pick_with_blacklist

import (
	"fmt"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	n         int
	blackList []int
}

type ans struct {
}

func Test(t *testing.T) {

	cases := []question{
		{param{7, []int{2, 3, 5}}, ans{}},
		//{param{3, []int{2}}, ans{}},
		//{param{4, []int{0, 3, 1}}, ans{}},
	}

	for _, v := range cases {
		c := Constructor(v.param.n, v.param.blackList)

		m := make(map[int]int)
		for i := 0; i < 1000; i++ {
			m[c.Pick()]++
		}

		fmt.Printf("%v\n", m)
	}

}
