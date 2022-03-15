package _025_reverse_k_group

import "tech.jiangchen/go-study/src/algs/data_structure"

type ListNode = data_structure.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		// 不足k个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前k个
	newHead := reverse(a, b)
	// 递归反转后续链表并连接起来
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 反转区间 [a, b) 的元素，注意是左闭右开
func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	for cur != b {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
