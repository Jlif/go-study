package _316_remove_duplicate_letters

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
	ans string
}

func Test(t *testing.T) {

	cases := []question{
		{param{"bcabc"}, ans{"abc"}},
	}

	for _, v := range cases {
		assert.New(t).Equal(v.ans.ans, removeDuplicateLetters(v.param.s))
	}

}
