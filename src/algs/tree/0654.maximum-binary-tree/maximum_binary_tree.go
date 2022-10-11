package _654_maximum_binary_tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal := -1
	index := -1
	for i, val := range nums {
		if val > maxVal {
			maxVal = val
			index = i
		}
	}

	root := TreeNode{Val: maxVal}
	root.Left = constructMaximumBinaryTree(nums[0:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return &root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
