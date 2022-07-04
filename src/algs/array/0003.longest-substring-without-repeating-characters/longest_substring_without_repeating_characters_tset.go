package _003_longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	s string
}

type ans struct {
	ans int
}

func Test(t *testing.T) {

	cases := []question{
		{param{"😁哈哈"}, ans{1}},
	}

	for _, v := range cases {
		assert.New(t).Equal(lengthOfLongestSubstring(v.param.s), v.ans.ans)
	}

}
