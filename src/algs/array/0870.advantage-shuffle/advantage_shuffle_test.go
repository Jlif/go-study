package _870_advantage_shuffle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	nums1 []int
	nums2 []int
}

type ans struct {
	ans []int
}

func Test(t *testing.T) {

	cases := []question{
		{param{[]int{12, 24, 8, 32}, []int{13, 25, 32, 11}}, ans{[]int{24, 32, 8, 12}}},
	}

	for _, v := range cases {
		assert.New(t).Equal(v.ans.ans, advantageCount(v.param.nums1, v.param.nums2))
	}
}
