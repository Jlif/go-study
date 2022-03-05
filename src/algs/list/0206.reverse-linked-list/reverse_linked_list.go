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

// 使用递归解法反转链表
func reverseLinkedList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
