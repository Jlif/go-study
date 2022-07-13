package _438_find_all_anagrams_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	s1 string
	s2 string
}

type ans struct {
	arr []int
}

func Test(t *testing.T) {

	cases := []question{
		{param{"cbaebabacd", "abc"}, ans{[]int{0, 6}}},
		{param{"abab", "ab"}, ans{[]int{0, 1, 2}}},
	}

	for _, v := range cases {
		assert.New(t).Equal(v.ans.arr, findAnagrams(v.param.s1, v.param.s2))
	}

}
