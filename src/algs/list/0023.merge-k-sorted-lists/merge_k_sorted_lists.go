package _023_merge_k_sorted_lists

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
)

type ListNode = data_structure.ListNode
type PQ = data_structure.PQ

func mergeKLists(lists []*ListNode) *ListNode {
	pq := PQ{}
	for i := 0; i < cap(lists); i++ {
		pq.Push(lists[i])
	}
	return nil
}
