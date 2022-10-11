package _654_maximum_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	param
	ans
}

type param struct {
	s []int
}

type ans struct {
	root *TreeNode
}

func Test(t *testing.T) {

	cases := []question{
		{param{[]int{3, 2, 1, 6, 0, 5}}, ans{nil}},
	}

	for _, v := range cases {
		assert.New(t).Equal(v.ans.root, constructMaximumBinaryTree(v.param.s))
	}

}
