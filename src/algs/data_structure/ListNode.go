package data_structure

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr []int) *ListNode {
	init := &ListNode{
		Val:  0,
		Next: nil,
	}
	tmp := init
	for i := 0; i < len(arr); i++ {
		tmp.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		tmp = tmp.Next
	}
	return init.Next
}

func PrintListNode(node *ListNode) {
	print(node.Val, " ")
	if node.Next != nil {
		PrintListNode(node.Next)
	}
	println()
}
