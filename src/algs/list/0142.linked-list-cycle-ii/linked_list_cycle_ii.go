package _142_linked_list_cycle_ii

import "tech.jiangchen/go-study/src/algs/data_structure"

type ListNode = data_structure.ListNode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	var flag bool
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			flag = true
			break
		}
	}
	if !flag {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
