package main

import "tech.jiangchen/go-study/src/algs/data_structure"

type ListNode = data_structure.ListNode

func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoSortedLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoSortedLists(l1, l2.Next)
	return l2
}
