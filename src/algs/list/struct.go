package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(arr []int) *ListNode {
	init := ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := init
	for i := 0; i < len(arr); i++ {
		tmp.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		tmp = *tmp.Next
	}
	return init.Next
}

func printListNode(node *ListNode) {
	print(node.Val, " ")
	if node.Next != nil {
		printListNode(node.Next)
	}
	println()
}
