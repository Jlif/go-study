package _187_repeated_dna_sequences

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
	ans []string
}

func Test(t *testing.T) {

	cases := []question{
		{param{"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"}, ans{[]string{"AAAAACCCCC", "CCCCCAAAAA"}}},
		{param{"AAAAAAAAAAAAA"}, ans{[]string{"AAAAAAAAAA"}}},
	}

	for _, v := range cases {
		assert.New(t).Equal(v.ans.ans, findRepeatedDnaSequences(v.param.s))
	}

}
