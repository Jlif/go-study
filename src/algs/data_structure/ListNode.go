package data_structure

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNodeFromArr(arr []int) *ListNode {
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

func GetListNodeStr(node *ListNode) string {
	if node == nil {
		return "{}"
	}
	str := "{" + strconv.Itoa(node.Val)
	node = node.Next
	count := 0
	for node != nil {
		str = str + "," + strconv.Itoa(node.Val)
		node = node.Next
		count++
		if count > 10000 {
			println("链表过长，考虑是否成环")
			break
		}
	}
	return str + "}"
}
