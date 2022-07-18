package _704_binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	nums   []int
	target int
}

type ans struct {
	ans int
}

func Test(t *testing.T) {

	cases := []question{
		{param{[]int{-1, 0, 3, 5, 9, 12}, 9}, ans{4}},
		{param{[]int{5}, 5}, ans{0}},
	}

	for _, v := range cases {
		assert.New(t).Equal(v.ans.ans, search(v.param.nums, v.param.target))
	}
}
