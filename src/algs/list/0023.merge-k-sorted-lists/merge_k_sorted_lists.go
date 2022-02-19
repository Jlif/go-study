package _023_merge_k_sorted_lists

import (
	"tech.jiangchen/go-study/src/algs/data_structure"
)

type ListNode = data_structure.ListNode

func mergeKSortedLists1(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return lists[0]
	}
	var tmp *ListNode
	for _, elem := range lists {
		tmp = mergeTwoSortedList(tmp, elem)
	}
	return tmp
}

// 方法二：使用归并算法
func mergeKSortedLists2(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return lists[0]
	}
	num := size / 2
	left := mergeKSortedLists2(lists[:num])
	right := mergeKSortedLists2(lists[num:])
	return mergeTwoSortedList(left, right)
}

func mergeTwoSortedList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoSortedList(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoSortedList(list1, list2.Next)
	return list2
}
