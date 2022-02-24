package _206_reverse_linked_list

import "tech.jiangchen/go-study/src/algs/data_structure"

type ListNode = data_structure.ListNode

func reverseLinkedList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
