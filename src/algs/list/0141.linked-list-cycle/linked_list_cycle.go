package _141_linked_list_cycle

import "tech.jiangchen/go-study/src/algs/data_structure"

type ListNode = data_structure.ListNode

func hasCycle(head *ListNode) bool {

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false

}
