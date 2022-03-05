package _092_reverse_linked_list_ii

import "tech.jiangchen/go-study/src/algs/data_structure"

type ListNode = data_structure.ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var successor *ListNode

// 反转以 head 为起点的 n 个节点，返回新的头结点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
