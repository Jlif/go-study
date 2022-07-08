package _567_permutation_in_string

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
	ans bool
}

func Test(t *testing.T) {

	cases := []question{
		{param{"ab", "eidbaooo"}, ans{true}},
		{param{"ab", "eidboaoo"}, ans{false}},
		{param{"hello", "ooolleoooleh"}, ans{false}},
		{param{"abcdxabcde", "abcdeabcdx"}, ans{true}},
	}

	for _, v := range cases {
		assert.New(t).Equal(v.ans.ans, checkInclusion(v.param.s1, v.param.s2))
	}

}
