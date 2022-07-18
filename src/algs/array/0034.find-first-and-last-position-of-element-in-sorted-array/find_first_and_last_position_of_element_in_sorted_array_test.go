package _034_find_first_and_last_position_of_element_in_sorted_array

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
	ans []int
}

func Test(t *testing.T) {

	cases := []question{
		{param{[]int{5, 7, 7, 8, 8, 10}, 8}, ans{[]int{3, 4}}},
	}

	for _, v := range cases {
		assert.New(t).Equal(v.ans.ans, searchRange(v.param.nums, v.param.target))
	}

}
