package _023_merge_k_sorted_lists

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
)

type ListNode = data_structure.ListNode
type PriorityQueue = data_structure.PriorityQueue

func mergeKSortedLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := ListNode{}
	p := dummy
	pq := PriorityQueue{}
	for i := 0; i < len(lists); i++ {
		pq.Push(lists[i])
	}
	for {
		if len(pq) != 0 {
			node := pq.Pop()
			p.Next := *node
			//if node.Next != nil {
			//	pq.Push(node.Next)
			//}
			//p = *p.Next
		}
	}
	return dummy.Next
}
